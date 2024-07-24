package main

import (
	"RestApi/controllers/handlers"
	"RestApi/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/wallets/users", handlers.HomeHandler)
	r.HandleFunc("/wallets/users/{id}", service.UserHandler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
