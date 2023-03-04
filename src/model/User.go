package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Nickname   string    `json:"nickname,omitempty"`
	SSN        string    `json:"ssn,omitempty"`
	Password   string    `json:"password,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}

func (u *User) ValidateFields() error {
	u.emptySpaces()
	if u.Name == "" {
		return errors.New("error: field name can't be a empty value")
	}
	if u.Nickname == "" {
		return errors.New("error: field nickname can't be a empty value")
	}
	if u.SSN == "" {
		return errors.New("error: field SSN can't be a empty value")
	}
	if u.Password == "" {
		return errors.New("error: field password can't be a empty value")
	}

	return nil
}

func (u *User) emptySpaces() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nickname = strings.TrimSpace(u.Nickname)
	u.SSN = strings.TrimSpace(u.SSN)
	u.Password = strings.TrimSpace(u.Password)
}
