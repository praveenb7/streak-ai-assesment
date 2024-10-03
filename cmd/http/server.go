package main

import (
	"log"
	"main/cmd/http/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "8090"

	r := mux.NewRouter()

	r.HandleFunc("/find-pairs", handlers.FindPairs)

	log.Println("Listening on port:", port)
	http.ListenAndServe(":"+port, r)
}
