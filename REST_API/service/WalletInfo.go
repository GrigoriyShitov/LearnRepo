package service

import (
	"RestApi/controllers/methods"
	"RestApi/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserInfo struct {
	User    storage.User
	Wallets []storage.Wallet
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var ReturnVal UserInfo

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	ReturnVal.User = methods.GetUserInfo(ctx, uint32(id))
	ReturnVal.Wallets = methods.GetWalletsInfo(ctx, uint32(id))

	data, err := json.MarshalIndent(ReturnVal, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	w.Write([]byte(data))
}
