package tool

import "golang.org/x/crypto/bcrypt"

func Encryption(password string) (hash []byte, err error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CompareHash(inPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(inPassword))
}
