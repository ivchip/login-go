package routers

import (
	"encoding/json"
	"github.com/ivchip/login-go/bd"
	"github.com/ivchip/login-go/jwt"
	"github.com/ivchip/login-go/models"
	"net/http"
	"time"
)

func Login(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set("Content-Type", "application/json")
	var t models.User
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Email or password invalid `+ err.Error() +`"}`))
		return
	}
	if len(t.Email) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Email is required"}`))
		return
	}
	document, exists := bd.TriedLogin(t.Email, t.Password)
	if !exists {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Email or password invalid"}`))
		return
	}
	jwtKey, err1 := jwt.GenerateJWT(document)
	if err1 != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"An error occurred while trying to generate the token `+ err1.Error() +`"}`))
		return
	}
	resp := models.ResponseLogin {
		Token: jwtKey,
	}
	response.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(response).Encode(resp)
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(response, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}