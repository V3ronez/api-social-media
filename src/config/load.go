package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Port = 0
var StringConnection = ""

func LoadDB() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error to load .env")
	}
	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_name"))
}
