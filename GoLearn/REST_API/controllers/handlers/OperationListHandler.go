package handlers

import (
	middleware "RestApi/controllers/midleware"
	"RestApi/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) OperationListHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := ctx.Value(middleware.UserId).(uint)
	WalletNum, err := strconv.ParseUint(mux.Vars(r)["WalletNum"], 10, 64)
	if err != nil {
		w.Write([]byte("Parse idWallet error: " + err.Error()))
		return
	}

	Operations, err := service.OperationList(ctx, id, uint(WalletNum))
	if err != nil {
		w.Write([]byte("Op error: " + err.Error()))
		return
	}

	fmt.Println(string(Operations))
	w.Write([]byte(Operations))
}
