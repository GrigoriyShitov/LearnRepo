package handlers

import (
	"RestApi/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) MakeOperationHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	idUser, err := strconv.ParseUint(mux.Vars(r)["idUser"], 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	idWallet, err := strconv.ParseUint(mux.Vars(r)["idWallet"], 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	Type := mux.Vars(r)["type"]
	if Type != "deposit" && Type != "withdraw" {
		w.Write([]byte("Wrong operation type"))
		return
	}
	Amount, err := strconv.ParseFloat(mux.Vars(r)["amount"], 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	Category := mux.Vars(r)["category"]
	data, err := service.NewOperationWithWallet(ctx, uint(idUser), uint(idWallet), Type, Amount, Category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}
