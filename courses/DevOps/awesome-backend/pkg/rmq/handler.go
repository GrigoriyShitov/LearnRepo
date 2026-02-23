package rmq

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"git.itmo.su/go-modules/log/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Handler interface {
	Ping(ctx context.Context) error
	Stop(ctx context.Context) error
	Start(ctx context.Context) error
	GetItems(ctx context.Context, queueName string) (interface{}, error)
	SendData(ctx context.Context, queueName string, body string) error
}

func InitHandler(user, password, host, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port) //nolint:nosprintfhostport
}

func NewRabbitMqHandler(dns string) Handler {
	if dns == "" {
		panic("nil amqp handler")
	}
	return &rabbitHandler{dsn: dns}
}

type rabbitHandler struct {
	dsn  string
	conn *amqp.Connection
}

func (r *rabbitHandler) Ping(ctx context.Context) error {
	url := fmt.Sprintf("%s/api/aliveness-test/%s", "http://rabbitmq:15672", "%2f")
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	req.SetBasicAuth("guest", "guest")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("http status %d", resp.StatusCode)
	}

	// Ответ должен быть {"status":"ok"}
	return nil
}

func (r *rabbitHandler) SendData(ctx context.Context, query string, body string) error {
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
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
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

func (r *rabbitHandler) GetItems(ctx context.Context, queueName string) (interface{}, error) {
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
		ctx,
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
		timeoutContext, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		select {
		case <-timeoutContext.Done():
			cancel()
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
