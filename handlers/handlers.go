package handlers

import (
	"github.com/gorilla/mux"
	"github.com/ivchip/login-go/middleware"
	"github.com/ivchip/login-go/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Managers()  {
	router := mux.NewRouter()
	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}