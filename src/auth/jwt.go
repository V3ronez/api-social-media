package auth

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

func Jwt(userId string) (string, error) {
	p := jwt.MapClaims{}
	p["userId"] = userId
	p["authorized"] = true
	p["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_TOKEN")))
	if err != nil {
		return "", err
	}
	return t, nil
}
