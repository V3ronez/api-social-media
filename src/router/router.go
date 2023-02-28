package router

import "github.com/gorilla/mux"

func NewRoute() *mux.Router {
	return mux.NewRouter()
}
