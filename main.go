package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadDB()
	fmt.Println("⚡️ loading all variables... ⚡️")

	database.ConnectDB()
	fmt.Printf("⚡️ server running on port %d... ⚡️\n", config.Port)
	r := router.NewRoute()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
