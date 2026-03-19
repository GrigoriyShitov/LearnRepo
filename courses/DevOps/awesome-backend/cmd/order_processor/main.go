package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Структуры данных
type Order struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type OrderMessage struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Глобальные переменные
var (
	dbPool      *pgxpool.Pool
	minioClient *minio.Client
	rabbitConn  *amqp.Connection
	rabbitCh    *amqp.Channel
	s3Bucket    string
)

// Инициализация подключений
func init() {
	// MinIO клиент
	endpoint := getEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKey := getEnv("MINIO_ACCESS_KEY", "minioadmin")
	secretKey := getEnv("MINIO_SECRET_KEY", "minioadmin")

	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}

	s3Bucket = getEnv("MINIO_BUCKET", "order-logs")

	// Создаем bucket если не существует
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, s3Bucket)
	if err != nil {
		log.Printf("Warning: could not check bucket existence: %v", err)
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, s3Bucket, minio.MakeBucketOptions{})
		if err != nil {
			log.Printf("Warning: could not create bucket: %v", err)
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Подключение к PostgreSQL
func connectDB() (*pgxpool.Pool, error) {
	host := getEnv("ORDERS_DB_HOST", "localhost")
	port := getEnv("ORDERS_DB_PORT", "5432")
	user := getEnv("ORDERS_DB_USER", "awesome-backend")
	password := getEnv("ORDERS_DB_PASSWORD", "postgres")
	dbname := getEnv("ORDERS_DB_NAME", "devopsdb")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return pool, nil
}

// Сохранение лога в S3
func saveLogToS3(orderID int, message string) {
	ctx := context.Background()
	dateStr := time.Now().Format("02-01-2006") // DD-MM-YYYY
	filename := fmt.Sprintf("orders-%s.log", dateStr)

	var existingContent []byte

	// Пытаемся получить существующий файл
	obj, err := minioClient.GetObject(ctx, s3Bucket, filename, minio.GetObjectOptions{})
	if err == nil {
		defer obj.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(obj)
		if err == nil {
			existingContent = buf.Bytes()
		}
	}

	// Формируем новую строку лога
	logLine := fmt.Sprintf("[%s] Order %d: %s\n",
		time.Now().Format("15:04:05"), orderID, message)

	// Объединяем существующий контент с новым
	newContent := append(existingContent, []byte(logLine)...)

	// Загружаем обратно в MinIO
	_, err = minioClient.PutObject(ctx, s3Bucket, filename,
		bytes.NewReader(newContent), int64(len(newContent)),
		minio.PutObjectOptions{ContentType: "text/plain"})

	if err != nil {
		log.Printf("S3 upload failed: %v", err)
	}
}

// Обработчик HTTP запросов
func getOrderHandler(c *gin.Context) {
	var orderID int
	_, err := fmt.Sscanf(c.Param("order_id"), "%d", &orderID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	var order Order
	query := "SELECT id, status, description FROM orders WHERE id = $1"
	err = dbPool.QueryRow(context.Background(), query, orderID).Scan(
		&order.ID, &order.Status, &order.Description)

	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(404, gin.H{"error": "Order not found"})
		} else {
			c.JSON(500, gin.H{"error": "Database error"})
		}
		return
	}

	// Логируем запрос статуса
	go saveLogToS3(orderID, fmt.Sprintf("Status requested - current status: %s", order.Status))

	c.JSON(200, order)
}

// Обработка сообщений из RabbitMQ
func processMessage(delivery amqp.Delivery) {
	defer delivery.Ack(false)

	var msg OrderMessage
	err := json.Unmarshal(delivery.Body, &msg)
	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	// Вставляем заказ в БД
	query := "INSERT INTO orders (id, status, description) VALUES ($1, 'created', $2)"
	_, err = dbPool.Exec(context.Background(), query, msg.ID, msg.Description)
	if err != nil {
		log.Printf("Failed to insert order: %v", err)
		return
	}

	// Логируем создание заказа
	saveLogToS3(msg.ID, fmt.Sprintf("Order %d created: %s", msg.ID, msg.Description))
	log.Printf("Processed order: %d", msg.ID)
}

// Consumer для RabbitMQ
func startConsumer() {
	var err error

	// Подключаемся к RabbitMQ
	rmqHost := getEnv("RMQ_HOST", "rabbitmq")
	rmqUser := getEnv("RMQ_USER", "guest")
	rmqPass := getEnv("RMQ_PASSWORD", "guest")

	rabbitConn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:5672/", rmqUser, rmqPass, rmqHost))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	rabbitCh, err = rabbitConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	queueName := getEnv("RMQ_QUEUE", "orders_queue")
	queue, err := rabbitCh.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// Подписываемся на очередь
	msgs, err := rabbitCh.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack (мы будем подтверждать вручную)
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("Processor started, waiting for messages...")

	// Бесконечно обрабатываем сообщения
	for msg := range msgs {
		processMessage(msg)
	}
}

func main() {
	// Подключаемся к БД
	var err error
	dbPool, err = connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	// Запускаем consumer в отдельной горутине
	go startConsumer()

	// Настраиваем HTTP сервер
	r := gin.Default()
	r.GET("/order/:order_id", getOrderHandler)

	// Graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Ждем сигнала завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Закрываем соединения
	if rabbitCh != nil {
		rabbitCh.Close()
	}
	if rabbitConn != nil {
		rabbitConn.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
