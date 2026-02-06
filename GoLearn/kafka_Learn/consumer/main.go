package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	// Настройка конфигурации Consumer
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Создание нового Consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Ошибка при создании Consumer: %v", err)
	}
	defer consumer.Close()

	// Подписка на topic
	partitionConsumer, err := consumer.ConsumePartition("study", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Ошибка при подписке на topic: %v", err)
	}
	defer partitionConsumer.Close()

	// Чтение сообщений
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Получено сообщение: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Ошибка: %v\n", err)
		}
	}
}
