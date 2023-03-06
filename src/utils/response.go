package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JsonResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}
