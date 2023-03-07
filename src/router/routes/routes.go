package routes

import (
	"api/src/router/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

// initialize all user routes
func InitRoutes(r *mux.Router) *mux.Router {
	routes := routesUser
	routes = append(routes, login)

	for _, route := range routes {
		if route.NeedAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Auth(route.Function)),
			).Methods(route.Method)
			// r.HandleFunc(route.URI, middlewares.Auth(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
