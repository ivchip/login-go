package routers

import (
	"encoding/json"
	"github.com/ivchip/login-go/bd"
	"github.com/ivchip/login-go/models"
	"net/http"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$").MatchString
var lowerCase = regexp.MustCompile("^[a-z]+$").MatchString
var upperCase = regexp.MustCompile("^[A-Z]+$").MatchString
var number = regexp.MustCompile("^[0-9]+$").MatchString
var symbol = regexp.MustCompile("[^A-Za-z0-9 ]").MatchString

func Register(response http.ResponseWriter, request *http.Request)  {
	response.Header().Set("Content-Type", "application/json")
	var t models.User
	if validateData(&t, response, request) {
		return
	}
	_, status, err := bd.InsertRegister(t)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"An error occurred while trying to register the user `+ err.Error() +`"}`))
		return
	}
	if !status {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Failed to insert user record"}`))
		return
	}
	response.WriteHeader(http.StatusCreated)
	_, _ = response.Write([]byte(`{"message":"Success"}`))
}

func validateData(t *models.User, response http.ResponseWriter, request *http.Request) bool {
	err := json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Error in request data ` + err.Error() + `"}`))
		return true
	}
	if len(t.Email) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Email address is required"}`))
		return true
	}
	if !isEmailValid(t.Email) {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Email address is invalid"}`))
		return true
	}
	if len(t.Password) < 8 {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Password must be greater than 8 characters"}`))
		return true
	}
	if !isPasswordSecure(t.Password) {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"Password is not secure. Password must contain uppercase, lowercase, numbers and special characters"}`))
		return true
	}
	_, exists, _ := bd.CheckIsExitsUser(t.Email)
	if exists {
		response.WriteHeader(http.StatusBadRequest)
		_, _ = response.Write([]byte(`{"error":"There is already a registered user with that email"}`))
		return true
	}
	return false
}

func isEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 100 {
		return false
	}
	return emailRegex(email)
}

func isPasswordSecure(password string) bool {
	hasLowercaseLetters := false
	hasUppercaseLetters := false
	hasNumbers := false
	hasSymbols := false
	for _, letter := range password {
		if lowerCase(string(letter)) {
			hasLowercaseLetters = true
			break
		}
	}
	for _, letter := range password {
		if upperCase(string(letter)) {
			hasUppercaseLetters = true
			break
		}
	}
	for _, letter := range password {
		if number(string(letter)) {
			hasNumbers = true
			break
		}
	}
	for _, letter := range password {
		if symbol(string(letter)) {
			hasSymbols = true
			break
		}
	}
	if hasLowercaseLetters && hasUppercaseLetters && hasNumbers && hasSymbols {
		return true
	}
	return false
}