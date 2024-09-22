package handlers

import (
	"log"
	"net/http"

	secret "RestApi/.env"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func (h *handler) AuthHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["idUser"]
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"iat": 1725287400,
		})
	s, err := t.SignedString([]byte(secret.SecretKey))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(s)
	w.Write([]byte("Your token is: " + s))

}
