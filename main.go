package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadDB()
	fmt.Println("⚡️ loading all variables... ⚡️")
	fmt.Printf("⚡️ server running on port %d... ⚡️", config.Port)
	r := router.NewRoute()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
