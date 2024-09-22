package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// данные пакеты нужны для системы проверки

func main() {
	// http запрос с методом GET
	var name string
	var age int
	fmt.Scan(&name, &age)
	baseUrl := "http://127.0.0.1:8080/hello"
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", fmt.Sprint(age))
	fullUrl := baseUrl + "?" + params.Encode()
	resp, err := http.Get(fullUrl)
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
