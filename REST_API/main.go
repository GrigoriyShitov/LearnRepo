package main

import (
	"RestApi/controllers/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/wallets/users", handlers.HomeHandler)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}", handlers.GetUserInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}", handlers.OperationListHandler).Methods(http.MethodGet)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{idOperation:[0-9]+}", handlers.OperationDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/wallets/users/{whoChange:[0-9]+}/{idUser:[0-9]+}/{role:admin|user}", handlers.ADMIN_EditRoleHandler).Methods(http.MethodPatch)
	r.HandleFunc("/wallets/users/{idUser:[0-9]+}/{idWallet:[0-9]+}/{type:deposit|withdraw}/{amount:[0-9]+}/{category}", handlers.MakeOperationHandler).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}
