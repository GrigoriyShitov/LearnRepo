package handlers

import (
	"RestApi/service"
	"RestApi/storage"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func MakeOperationHandler(w http.ResponseWriter, r *http.Request) {
	var (
		AddOperation storage.NewOperation
	)
	ctx := r.Context()
	Id, _ := strconv.ParseUint(mux.Vars(r)["idUser"], 10, 64)

	AddOperation.UserId = uint(Id)
	idWallet, err := strconv.ParseUint(mux.Vars(r)["idWallet"], 10, 64)
	if err != nil {
		w.Write([]byte("Hello, World!"))
		return
	}
	AddOperation.WalletId = uint(idWallet)
	AddOperation.Type = mux.Vars(r)["type"]
	if AddOperation.Type != "deposit" && AddOperation.Type != "withdraw" {
		w.Write([]byte("Hello, World!"))
		return
	}
	AddOperation.Amount, _ = strconv.ParseFloat(mux.Vars(r)["amount"], 64)
	AddOperation.Category = mux.Vars(r)["category"]
	err = service.NewOperationWithWallet(ctx, &AddOperation)
	if err != nil {
		w.Write([]byte("Hello, World!"))
		return
	}
}
