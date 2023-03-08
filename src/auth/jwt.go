package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

func Jwt(userId string) (string, error) {
	p := jwt.MapClaims{}
	p["userId"] = userId
	p["authorized"] = true
	p["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)

	t, err := token.SignedString(config.SecretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

func ValidateToken(r *http.Request) error {
	th := GetToken(r)
	// fmt.Println(th)

	tk, err := jwt.Parse(th, validateKeyToken)
	// fmt.Println(tk)
	if err != nil {
		return err
	}

	if _, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		return nil
	}
	return errors.New("token invalid")
}

func GetToken(r *http.Request) string {
	t := r.Header.Get("Authorization")
	if len(strings.Split(t, " ")) == 2 {
		return strings.Split(t, " ")[1]
	}
	return ""
}

func validateKeyToken(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method of signature invalid! %v", t.Header["alg"])
	}

	return config.SecretKey, nil
}
