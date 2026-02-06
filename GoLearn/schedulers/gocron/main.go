package main

import (
	"github.com/go-co-op/gocron"
	"time"
)

func task(task string) {
	println(task)
}

func main() {
	// инициализируем объект планировщика
	s := gocron.NewScheduler(time.UTC)
	// добавляем одну задачу на каждую минуту
	s.Every(5).Seconds().Do(task, "task")
	// запускаем планировщик с блокировкой текущего потока
	s.StartAsync()
}
