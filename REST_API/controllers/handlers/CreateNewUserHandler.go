package handlers

import (
	"RestApi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	Id, _ := strconv.Atoi(mux.Vars(r)["idUser"])
	name := mux.Vars(r)["name"]
	err := service.CreateNewUser(ctx, uint(Id),name)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

}
