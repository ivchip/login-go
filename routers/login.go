package routers

import (
	"encoding/json"
	"github.com/ivchip/login-go/bd"
	"github.com/ivchip/login-go/jwt"
	"github.com/ivchip/login-go/models"
	"net/http"
	"time"
)

func Login(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Add("content-type", "application/json")
	var t models.User
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Email or password invalid " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer, "Email is required", http.StatusBadRequest)
		return
	}
	document, exists := bd.TriedLogin(t.Email, t.Password)
	if !exists {
		http.Error(writer, "Email or password invalid", http.StatusBadRequest)
		return
	}
	jwtKey, err1 := jwt.GenerateJWT(document)
	if err1 != nil {
		http.Error(writer, "An error occurred while trying to generate the token " + err1.Error(), http.StatusBadRequest)
		return
	}
	resp := models.ResponseLogin {
		Token: jwtKey,
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}