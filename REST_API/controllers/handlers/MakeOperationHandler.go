package handlers

import (
	middleware "RestApi/controllers/midleware"
	"RestApi/service"
	"net/http"
	"strconv"
)

func (h *handler) MakeOperationHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userID := ctx.Value(middleware.UserId).(uint)

	WalletNum, err := strconv.ParseUint(r.FormValue("WalletNum"), 10, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	Type := r.FormValue("Type")
	if Type != "deposit" && Type != "withdraw" {
		w.Write([]byte("Wrong operation type"))
		return
	}
	Amount, err := strconv.ParseFloat(r.FormValue("Amount"), 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	Category := r.FormValue("Category")
	data, err := service.NewOperationWithWallet(ctx, userID, uint(WalletNum), Type, Amount, Category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}
