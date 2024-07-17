package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int = 0

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte(strconv.Itoa(counter)))
	case http.MethodPost:
		r.ParseForm()
		digit, err := strconv.Atoi(r.FormValue("count"))
		if err != nil {
			w.Write([]byte("это не число"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		counter += digit
		w.Write([]byte("OK"))
	}
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/count", handler)
	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
