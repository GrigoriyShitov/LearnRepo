package main

import (
	"RestApi/controllers/handlers"
	"log"
	"net/http"
)

func main() {

	h := handlers.HandlerInit()
	h.Router.HandleFunc("/wallets/users", h.HomeHandler)
	h.Router.HandleFunc("/wallets/users/{idUser:[0-9]+}", h.GetUserInfoHandler).Methods(http.MethodGet)
	h.Router.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}", h.OperationListHandler).Methods(http.MethodGet)
	h.Router.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{idOperation:[0-9]+}", h.OperationDeleteHandler).Methods(http.MethodDelete)
	h.Router.HandleFunc("/wallets/users/{whoChange:[0-9]+}/{idUser:[0-9]+}/{role:admin|user}", h.ADMIN_EditRoleHandler).Methods(http.MethodPatch)
	h.Router.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{type:deposit|withdraw}/{amount:[0-9]+}/{category}", h.MakeOperationHandler).Methods(http.MethodPost)
	log.Println("API is running!")
	http.ListenAndServe(":8080", h.Router)
}
