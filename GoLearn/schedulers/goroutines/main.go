package main

import (
	"context"
	"time"
)

func task(ctx context.Context) {
	// запускаем бесконечный цикл
	for {
		select {
		// проверяем не завершён ли ещё контекст и выходим, если завершён
		case <-ctx.Done():
			return

		// выполняем нужный нам код
		default:
			println("Hello gophers!")
		}
		// делаем паузу перед следующей итерацией
		time.Sleep(time.Minute)
	}
}

func main() {
	// создаём контекст с функцией завершения
	ctx, cancel := context.WithCancel(context.Background())
	// запускаем нашу горутину
	go task(ctx)
	// делаем паузу, чтобы дать горутине поработать
	time.Sleep(10 * time.Second)
	// завершаем контекст, чтобы завершить горутину
	cancel()
}
