package main

import (
	"github.com/ivchip/login-go/bd"
	"github.com/ivchip/login-go/handlers"
	"log"
)

func main()  {
	if bd.CheckConnection() == 0 {
		log.Fatal("Without connection to DB")
		return
	}
	handlers.Managers()
}