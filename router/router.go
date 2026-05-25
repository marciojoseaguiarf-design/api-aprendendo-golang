package router

import (
	"api-aula-1/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {

	rotas := mux.NewRouter()

	routes.Register(rotas)

	return rotas
}

