package handlers

import (
	middleware "RestApi/controllers/midleware"
	"RestApi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) ADMIN_EditRoleHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	whoChange, ok := ctx.Value(middleware.UserId).(uint)
	if !ok {
		log.Println("sueta")
		return
	}
	Id, err := strconv.ParseUint(mux.Vars(r)["idUser"], 10, 64)
	if err != nil {
		w.Write([]byte("Parse Id error: " + err.Error()))
		return
	}

	role := mux.Vars(r)["role"]
	err = service.EditRoleService(ctx, uint(whoChange), uint(Id), role)
	if err != nil {
		w.Write([]byte("EditRole error: " + err.Error()))
		return
	}
	w.Write([]byte("OK"))
}
