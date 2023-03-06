package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 8)
}
func ComparePasswordAndHash(password, passwordWithHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordWithHash))
}
