package handlers

import (
	"RestApi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) CreateNewWalletHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	Id, _ := strconv.Atoi(mux.Vars(r)["idUser"])

	err := service.CreateNewWallet(ctx, uint(Id))

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
}
