package middlewares

import (
	"fmt"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authenticating...")
		next(w, r)
	}
}
