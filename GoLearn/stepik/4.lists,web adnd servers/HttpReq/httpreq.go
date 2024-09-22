package main

import (
	"fmt"
	"io"
	"net/http"
)

// данные пакеты нужны для системы проверки

func main() {
	// http запрос с методом GET
	resp, err := http.Get("http://127.0.0.1:5555/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close() // закрываем тело ответа после работы с ним

	data, err := io.ReadAll(resp.Body) // читаем ответ
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) // печатаем ответ как строку
}
