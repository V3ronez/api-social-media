package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUser = []Route{
	{
		URI:                "/user/{id}",
		Method:             http.MethodGet,
		Function:           controllers.GetUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Function:           controllers.GetAllUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/user",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/user/{id}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/user{id}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: false,
	},
}
