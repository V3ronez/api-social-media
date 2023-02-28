package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func NewRoute() *mux.Router {
	r := mux.NewRouter()
	return routes.InitRoutes(r)
}
