package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("⚡️ server running on port 8000... ⚡️")
	r := router.NewRoute()
	log.Fatal(http.ListenAndServe(":8000", r))
}
