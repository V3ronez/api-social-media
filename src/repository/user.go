package repository

import (
	"api/src/model"
	"database/sql"
	"fmt"
	"strings"
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
func (repository Users) Search(param string) ([]model.User, error) {
	p := strings.TrimSpace(param)
	query := fmt.Sprintf(`SELECT name, nickname, ssn, created_at FROM users WHERE name LIKE '%%%s%%'`, p)
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var us []model.User
	for rows.Next() {
		var u model.User
		if err = rows.Scan(&u.Name, &u.Nickname, &u.SSN, &u.Created_at); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}
