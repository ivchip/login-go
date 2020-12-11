package bd

import (
	"github.com/ivchip/login-go/models"
	"golang.org/x/crypto/bcrypt"
)

func TriedLogin(email string, password string) (models.User, bool) {
	u, exists, _ := CheckIsExitsUser(email)
	if !exists {
		return u, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return u, false
	}
	return u, true
}