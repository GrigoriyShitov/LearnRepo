package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// User - структура для представления объекта пользователя
type User struct {
	Name string `json:"name"`
	ID   uint32 `json:"id"`
}

// Output - структура для представления ответа сервера
type Output struct {
	JSON struct {
		Name string `json:"name"`
		ID   uint32 `json:"id"`
	} `json:"json"`
	URL string `json:"url"`
}

func main() {
	// Создаем экземпляр структуры User
	var u = User{
		Name: "Alex",
		ID:   1,
	}

	// Кодируем структуру User в JSON
	bytesRepresentation, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	// Отправляем POST-запрос на сервер с JSON-телом
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	// Читаем и конвертируем тело ответа в байты
	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Создаем экземпляр ответа сервера
	var out Output
	// Декодируeм данные в формате JSON и заполняем структуру
	err = json.Unmarshal(bytesResp, &out)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", out) // печатаем ответ в виде структуры
	fmt.Println(out.URL)     // печатаем конкретное поле структуры
}
