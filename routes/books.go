package routes

import (
	"api-aula-1/controller"
	"net/http"
)

var booksRoutes = []Route{
	{
		Url:    "/books",
		Method: http.MethodGet,
		Func:   controller.HandleSearch,
	},
}