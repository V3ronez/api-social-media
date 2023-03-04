package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"io"
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
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.ValidateFields(); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repository.InitRepository(db)
	userId, err := repository.CreateUser(user)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	userCreated := model.User{
		ID:       userId,
		Name:     user.Name,
		Nickname: user.Nickname,
		SSN:      user.SSN,
	}

	utils.JsonResponse(w, http.StatusCreated, userCreated)
	// w.Write([]byte(fmt.Sprintf("User create successfully! ID: %s", id))) // example to make a return
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user!"))
}
