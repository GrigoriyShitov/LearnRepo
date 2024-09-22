package handlers

import (
	middleware "RestApi/controllers/midleware"
	"RestApi/service"
	"log"
	"net/http"
)

func (h *handler) CreateNewWalletHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id := ctx.Value(middleware.UserId).(uint)
	err := service.CreateNewWallet(ctx, id)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
}
