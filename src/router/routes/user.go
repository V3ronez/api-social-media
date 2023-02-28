package routes

import "net/http"

var routes = []Route{
	{
		URI:                "/user",
		Method:             http.MethodGet,
		Function:           func(w http.ResponseWriter, r *http.Request) {},
		NeedAuthentication: false,
	},
	{
		URI:                "/user/{id}",
		Method:             http.MethodGet,
		Function:           func(w http.ResponseWriter, r *http.Request) {},
		NeedAuthentication: false,
	},
	{
		URI:                "/user",
		Method:             http.MethodPost,
		Function:           func(w http.ResponseWriter, r *http.Request) {},
		NeedAuthentication: false,
	},
	{
		URI:                "/user/{id}",
		Method:             http.MethodPut,
		Function:           func(w http.ResponseWriter, r *http.Request) {},
		NeedAuthentication: false,
	},
	{
		URI:                "/user{id}",
		Method:             http.MethodDelete,
		Function:           func(w http.ResponseWriter, r *http.Request) {},
		NeedAuthentication: false,
	},
}
