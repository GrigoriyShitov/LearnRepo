package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	// Настройка конфигурации Producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Создание нового Producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Ошибка при создании Producer: %v", err)
	}
	defer producer.Close()

	// Создание сообщения
	msg := &sarama.ProducerMessage{
		Topic: "study",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// Отправка сообщения
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Ошибка при отправке сообщения: %v", err)
	}

	fmt.Printf("Сообщение отправлено в partition %d с offset %d\n", partition, offset)
}
