package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 8)
}
func ComparePasswordAndHash(passwordWithHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(password))
}
