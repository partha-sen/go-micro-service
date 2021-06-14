package main

import (
	"log"
	"net/http"

	"learning.go/appservice/controller"
	"learning.go/appservice/db"
	"learning.go/appservice/middleware"
)

func addRouter() {

	http.HandleFunc("/openings", middleware.ValidateTokenAdminAccess(controller.HandleOpening))
	http.HandleFunc("/openings/", middleware.ValidateTokenAdminAccess(controller.HandleOpening))

	http.HandleFunc("/interviews", middleware.ValidateTokenUserAccess(controller.HandleInterview))
	http.HandleFunc("/interviews/", middleware.ValidateTokenUserAccess(controller.HandleInterview))

	http.HandleFunc("/questions", middleware.ValidateTokenUserAccess(controller.HandleQuestion))
	http.HandleFunc("/questions/", middleware.ValidateTokenUserAccess(controller.HandleQuestion))
}

func main() {
	db.SetupDatabase()
	addRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
