package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all user!"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error to get the body of request")
	}

	var user model.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal("Error to parser body json")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.InitRepository(db)
	repository.CreateUser(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}
