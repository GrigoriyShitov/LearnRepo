package handlers

import (
	middleware "RestApi/controllers/midleware"
	"RestApi/service"
	"fmt"
	"log"
	"net/http"
)

func (h *handler) GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id, ok := ctx.Value(middleware.UserId).(uint)
	if !ok {
		log.Println("sueta")
		return
	}
	data, err := service.UserWalletInfo(ctx, uint(id))
	if err != nil {
		w.Write([]byte(err.Error()))
		if data == nil {
			return
		}
	}
	fmt.Println(string(data))
	w.Write([]byte(data))
}
