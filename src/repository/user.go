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

func (repository Users) CreateUser(user model.User) (string, error) {
	defer repository.db.Close()

	var query = `INSERT INTO users (name, nickname, ssn, password) VALUES ($1, $2, $3, $4) RETURNING id`
	err := repository.db.QueryRow(query, &user.Name, &user.Nickname, &user.SSN, &user.Password).Scan(&user.ID)

	if err != nil {
		return "", err
	}
	return user.ID, err
}
