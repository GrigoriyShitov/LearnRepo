package handlers

import (
	"RestApi/service"
	"RestApi/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	var (
		ReturnVal *storage.UserInfo
		err       error
	)
	ctx := r.Context()
	Id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ReturnVal, err = service.UserWalletInfo(ctx, uint(Id))
	if err != nil {
		if err.Error() == "user not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		} else if err.Error() == "no wallets" {
			w.Write([]byte(err.Error()))
		}

	}
	data, err := json.MarshalIndent(ReturnVal, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	w.Write([]byte(data))
}