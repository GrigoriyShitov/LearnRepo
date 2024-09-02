package handlers

import (
	middleware "RestApi/controllers/midleware"
	"RestApi/service"
	"net/http"
	"strconv"
)

func (h *handler) OperationDeleteHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := ctx.Value(middleware.UserId).(uint)
	// idWallet, err := strconv.ParseUint(mux.Vars(r)["idWallet"], 10, 64)
	// if err != nil {
	// 	w.Write([]byte("Parse idWallet error: " + err.Error()))
	// 	return
	// }

	// idOperation, err := strconv.ParseUint(mux.Vars(r)["idOperation"], 10, 64)
	// if err != nil {
	// 	w.Write([]byte("Parse idOperation error: " + err.Error()))
	// 	return
	// }

	idWallet, err := strconv.ParseUint(r.FormValue("idWallet"), 10, 64)
	if err != nil {
		w.Write([]byte("Parse idWallet error: " + err.Error()))
		return
	}
	idOperation, err := strconv.ParseUint(r.FormValue("idOperation"), 10, 64)
	if err != nil {
		w.Write([]byte("Parse idOperation error: " + err.Error()))
		return
	}
	err = service.DeleteOperation(ctx, id, uint(idWallet), uint(idOperation))
	if err != nil {
		w.Write([]byte("Delete error: " + err.Error()))
		return
	}
	w.Write([]byte("deleted operation" + strconv.Itoa(int(idOperation)) + "from wallet" + strconv.Itoa(int(idWallet))))
}
