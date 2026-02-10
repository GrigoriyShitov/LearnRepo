package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message":"Hello, world"}`)
	})
	addr := ":8080"
	log.Printf("Go server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
