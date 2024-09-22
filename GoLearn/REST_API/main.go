package main

import (
	"RestApi/controllers/handlers"
	middleware "RestApi/controllers/midleware"
	"log"
	"net/http"
)

func main() {

	h := handlers.HandlerInit()

	h.Router.HandleFunc("/wallets/auth/{idUser:[0-9]+}", h.AuthHandler)
	h.Router.HandleFunc("/users/LK", http.HandlerFunc(h.GetUserInfoHandler)).Methods(http.MethodGet)                       // home page
	h.Router.Handle("/users/wallets/{WalletNum:[0-9]+}", http.HandlerFunc(h.OperationListHandler)).Methods(http.MethodGet) //get operations list                                         // get user info
	h.Router.Handle("/users/CreateNew", http.HandlerFunc(h.CreateNewUserHandler)).Methods(http.MethodPost)                 // create new user
	h.Router.Handle("/users/wallets/CreateNew", http.HandlerFunc(h.CreateNewWalletHandler)).Methods(http.MethodPost)       // create new wallet
	h.Router.Handle("/users/wallets/delete/{idWallet:[0-9]+}/{idOperation:[0-9]+}", http.HandlerFunc(h.OperationDeleteHandler)).Methods(http.MethodDelete)
	h.Router.Handle("/users/ChangeRole/{idUser:[0-9]+}/{role:admin|user}", http.HandlerFunc(h.ADMIN_EditRoleHandler)).Methods(http.MethodPatch)
	h.Router.Handle("/users/wallets", http.HandlerFunc(h.MakeOperationHandler)).Methods(http.MethodPost)
	h.Router.Use(middleware.JWTmiddleware)
	log.Println("API is running!")
	http.ListenAndServe(":8080", h.Router)
}
