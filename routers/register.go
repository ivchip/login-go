package routers

import (
	"encoding/json"
	"github.com/ivchip/login-go/bd"
	"github.com/ivchip/login-go/models"
	"net/http"
)

func Register(writer http.ResponseWriter, request *http.Request)  {
	var t models.User
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		http.Error(writer, "Error in request data " + err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(writer, "Email is required", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 8 {
		http.Error(writer, "Password must be greater than 8 characters", http.StatusBadRequest)
		return
	}
	_, exists, _ := bd.CheckIsExitsUser(t.Email)
	if exists {
		http.Error(writer, "There is already a registered user with that email", http.StatusBadRequest)
		return
	}
	_, status, err1 := bd.InsertRegister(t)
	if err1 != nil {
		http.Error(writer, "An error occurred while trying to register the user " + err1.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(writer, "Failed to insert user record", http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}