package handlers

import (
	"RestApi/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) OperationListHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	Id, err := strconv.ParseUint(mux.Vars(r)["idUser"], 10, 64)
	if err != nil {
		w.Write([]byte("Parse Id error: " + err.Error()))
		return
	}
	idWallet, err := strconv.ParseUint(mux.Vars(r)["idWallet"], 10, 64)
	if err != nil {
		w.Write([]byte("Parse idWallet error: " + err.Error()))
		return
	}

	Operations, err := service.OperationList(ctx, uint(Id), uint(idWallet))
	if err != nil {
		w.Write([]byte("Op error: " + err.Error()))
		return
	}
	data, err := json.MarshalIndent(Operations, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	w.Write([]byte(data))
}
