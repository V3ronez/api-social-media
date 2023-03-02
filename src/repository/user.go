package repository

import (
	"api/src/model"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) CreateUser(user model.User) (string, error) {
	return "", nil
}
