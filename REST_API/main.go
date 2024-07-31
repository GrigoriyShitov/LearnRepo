package main

import (
	"RestApi/controllers/handlers"
	"RestApi/storage/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	DB := db.Init()
	h := handlers.New(DB)
	r.HandleFunc("/wallets/users", h.HomeHandler)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}", h.GetUserInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}", h.OperationListHandler).Methods(http.MethodGet)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{idOperation:[0-9]+}", h.OperationDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/wallets/users/{whoChange:[0-9]+}/{idUser:[0-9]+}/{role:admin|user}", h.ADMIN_EditRoleHandler).Methods(http.MethodPatch)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{type:deposit|withdraw}/{amount:[0-9]+}/{category}", h.MakeOperationHandler).Methods(http.MethodPost)
	log.Println("API is running!")
	http.ListenAndServe(":8080", r)
}
