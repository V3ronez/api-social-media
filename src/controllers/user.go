package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id := p["id"]
	db, err := database.ConnectDB()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.InitRepository(db)
	u, err := repo.SearchForID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JsonResponse(w, http.StatusOK, u)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	param := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.ConnectDB()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.InitRepository(db)

	us, err := repo.Search(param)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, us)
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

	if err = user.ValidateFields(true); err != nil {
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

	p := mux.Vars(r)
	b, err := io.ReadAll(r.Body)

	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(b, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	uid, err := auth.GetUserId(r)
	if err != nil {
		utils.Error(w, http.StatusForbidden, err)
		return
	}

	if err = user.ValidateFields(false); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if p["id"] != uid {
		utils.Error(w, http.StatusForbidden, errors.New("action not valid fot this user"))
		return
	}

	db, err := database.ConnectDB()

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	u := repository.InitRepository(db)
	if err = u.UpdateUser(p["id"], user); err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JsonResponse(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)

	ut, err := auth.GetUserId(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	if p["id"] != ut {
		utils.Error(w, http.StatusForbidden, errors.New("impossible delete another user"))
		return
	}
	db, err := database.ConnectDB()

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	u := repository.InitRepository(db)

	if err = u.DeleteUser(p["id"]); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	utils.JsonResponse(w, http.StatusOK, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	s := p["userId"]
	uId, err := auth.GetUserId(r)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	if s == uId {
		utils.Error(w, http.StatusForbidden, errors.New("user can`t follower himself"))
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	ru := repository.InitRepository(db)
	err = ru.FollowUser(&uId, &s)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	utils.JsonResponse(w, http.StatusNoContent, nil)
}
