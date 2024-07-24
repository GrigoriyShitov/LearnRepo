package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello, World!"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
