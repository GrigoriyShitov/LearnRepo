package handlers

import (
	"RestApi/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	Id, _ := strconv.Atoi(mux.Vars(r)["idUser"])
	data, err := service.UserWalletInfo(ctx, uint(Id))
	if err != nil {
		if err.Error() == "user not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		} else if err.Error() == "no wallets" {
			w.Write([]byte(err.Error()))
		}

	}
	fmt.Println(string(data))
	w.Write([]byte(data))
}
