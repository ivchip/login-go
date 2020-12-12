package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ivchip/login-go/models"
	"github.com/ivchip/login-go/utils"
	"time"
)

func GenerateJWT(t models.User) (string, error) {
	secretKey, _ := utils.GetValueEnvironment("JWT.KEY")
	myKey := []byte(secretKey)
	payload := jwt.MapClaims{
		"email":t.Email,
		"firstName":t.FirstName,
		"lastName":t.LastName,
		"birthDate":t.BirthDate,
		"_id":t.ID.Hex(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}