package rmq

import (
	"context"
	"fmt"
	"time"

	"git.itmo.su/go-modules/log/v4"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitHandler interface {
	GetItems(ctx *fiber.Ctx, name string) (interface{}, error)
	Stop(ctx context.Context) error
	Start(ctx context.Context) error
	SendData(c *fiber.Ctx, query string, body string) error
}

func InitHandler(user, password, host, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port) //nolint:nosprintfhostport
}

func NewRabbitMqHandler(dns string) RabbitHandler {
	if dns == "nil" {
		panic("nil amqp handler")
	}
	return &rabbitHandler{dsn: dns}
}

type rabbitHandler struct {
	dsn  string
	conn *amqp.Connection
}

func (r *rabbitHandler) SendData(c *fiber.Ctx, query string, body string) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		query, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitHandler) GetItems(ctx *fiber.Ctx, queueName string) (interface{}, error) {
	ch, err := r.conn.Channel()
	defer ch.Close()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		queueName, // queueName
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	msgs, err := ch.ConsumeWithContext(
		ctx.Context(),
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		true,   // no-wait
		nil,    // args
	)
	var data []string
	go func() {
		for d := range msgs {
			data = append(data, string(d.Body))
		}
	}()
	for {
		timeoutContext, cancel := context.WithTimeout(ctx.Context(), 500*time.Millisecond)
		select {
		case <-timeoutContext.Done():
			return data, nil
		case d := <-msgs:
			data = append(data, string(d.Body))
			cancel()
			continue
		}
	}
}

func (r *rabbitHandler) Start(ctx context.Context) error {
	log.Infof("start RabbitMQ handler on %s", r.dsn)
	conn, err := amqp.Dial(r.dsn)
	if err != nil {
		return err
	}
	r.conn = conn
	return nil
}

func (r *rabbitHandler) Stop(ctx context.Context) error {
	return r.conn.Close()
}
