package middlewares

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//create a file before
		fmt.Printf("\n%s %s %s", r.Method, r.URL, r.Host)
		next(w, r)
	}
}
