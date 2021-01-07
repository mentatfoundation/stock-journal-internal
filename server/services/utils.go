package services

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordByte := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordByte, 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func PasswordsMatch(hashedPwd string, password string) bool {
	hashByte := []byte(hashedPwd)
	passwordByte := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashByte, passwordByte)
	if err != nil {
		return false
	}
	return true
}
