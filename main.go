package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"api-aula-1/routes"
)

func main() {
	r := mux.NewRouter()

	routes.Register(r)
	const addr = ":8080"

	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
