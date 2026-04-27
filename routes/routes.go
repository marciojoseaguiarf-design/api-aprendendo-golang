package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"api-aula-1/handler"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)
}