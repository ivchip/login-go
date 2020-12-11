package middleware

import (
	"github.com/ivchip/login-go/bd"
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(writer, "Lost connection to the DB", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(writer, request)
	}
}