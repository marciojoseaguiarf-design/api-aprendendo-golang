package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func Register(r *mux.Router) {
	var routes []Route
	routes = append(routes, usersRoutes...)
	routes = append(routes, booksRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.Url, route.Func).Methods(route.Method)
	}
}