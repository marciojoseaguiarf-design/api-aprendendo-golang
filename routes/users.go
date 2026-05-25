package routes

import (
	"api-aula-1/controller"
	"net/http"
)

var usersRoutes = []Route{
	{
		Url:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
	},
	{
		Url:    "/users",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
	},
	{
		Url:    "/users",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
	},
	{
		Url:    "/users",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
	},
}