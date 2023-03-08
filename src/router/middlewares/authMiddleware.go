package middlewares

import (
	"api/src/auth"
	"api/src/utils"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			utils.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
