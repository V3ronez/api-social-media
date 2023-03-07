package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/model"
	"api/src/repository"
	"api/src/security"
	"api/src/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var ul model.User
	if err = json.Unmarshal(b, &ul); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	repo := repository.InitRepository(db)
	ub, err := repo.SearchSSN(ul.SSN)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}

	if err = security.ComparePasswordAndHash(ub.Password, ul.Password); err != nil {
		err = errors.New("password or SSN is invalid")
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}
	t, err := auth.Jwt(ub.ID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(t))
}
