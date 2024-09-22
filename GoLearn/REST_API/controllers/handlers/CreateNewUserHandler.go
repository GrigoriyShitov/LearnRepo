package handlers

import (
	"RestApi/service"
	"log"
	"net/http"
)

func (h *handler) CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	// id, ok := ctx.Value(middleware.UserId).(uint)
	// if !ok {
	// 	log.Println(ok)
	// 	return
	// }
	name := r.FormValue("name")
	log.Println(name)

	err := service.CreateNewUser(ctx, name)

	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

}
