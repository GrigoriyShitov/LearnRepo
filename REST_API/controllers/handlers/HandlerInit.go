package handlers

import (
	"RestApi/service"

	"github.com/gorilla/mux"
)

type handler struct {
	Router *mux.Router
}

func HandlerInit() handler {
	r := mux.NewRouter()
	service.ServiceInit()
	var ret handler
	ret.Router = r
	return ret
}
