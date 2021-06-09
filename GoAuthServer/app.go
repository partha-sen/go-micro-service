package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"learning.go/authserver/config"
	"learning.go/authserver/controller"
	"learning.go/authserver/db"
	"learning.go/authserver/middleware"
	"learning.go/authserver/token"
)

func main() {
	db.SetupDatabase()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = config.DEFAULT_PORT
	}

	token.GlobalTokenStore = make(token.TokenStore)
	go token.RemoveExpiredToken()

	http.HandleFunc("/login", controller.HandleLogin)
	http.HandleFunc("/token/status/", controller.HandleTokenValidity)
	http.Handle("/logout", middleware.ValidateTokeMiddleware(controller.HandleLogOut))

	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
