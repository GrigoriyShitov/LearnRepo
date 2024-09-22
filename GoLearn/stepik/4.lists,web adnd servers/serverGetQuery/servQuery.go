package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RawQuery: ", r.URL.String())           // URL с параметрами
	fmt.Println("Method: ", r.Method)                   // тип метода
	fmt.Println("Name: ", r.URL.Query().Get("name"))    // значение параметра
	fmt.Println("IsExist: ", r.URL.Query().Has("name")) // существует ли такой параметр
	w.Write([]byte("Привет, Stepik!"))
}
func handler_api(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RawQuery: ", r.URL.String())           // URL с параметрами
	fmt.Println("Method: ", r.Method)                   // тип метода
	fmt.Println("Name: ", r.URL.Query().Get("name"))    // значение параметра
	fmt.Println("IsExist: ", r.URL.Query().Has("name")) // существует ли такой параметр
	w.Write([]byte(fmt.Sprintf("Hello,%s!", r.URL.Query().Get("name"))))
}

func handler_post(w http.ResponseWriter, r *http.Request) {
	// проверяем что метод POST
	if r.Method == "POST" {
		// читаем входящее тело запроса
		bytesBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Плохое тело запроса"))
			return
		}
		// печатаем тело запроса как строку
		fmt.Println(string(bytesBody))
		// отвечаем клиенту, что все хорошо
		w.Write([]byte("OK!"))
		return
	}
	w.Write([]byte("Разрешен только метод POST!"))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/user", handler_api)
	http.HandleFunc("/body", handler_post)
	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
