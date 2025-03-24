package internal

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return hashed, err
}

func VerifyPassword(hashed []byte, pass string) bool {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(pass))
	return err == nil
}
