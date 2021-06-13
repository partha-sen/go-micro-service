package main

import (
	"log"
	"net/http"

	"learning.go/appservice/controller"
	"learning.go/appservice/db"
	"learning.go/appservice/middleware"
)

func addRouter() {

	openingHandler := middleware.ValidateTokeMiddleware(
		middleware.ValidateAdminMiddleware(
			controller.HandleOpening))

	http.HandleFunc("/openings", openingHandler)
	http.HandleFunc("/openings/", openingHandler)

	interviewHandler := middleware.ValidateTokeMiddleware(
		middleware.ValidateUserMiddleware(
			controller.HandleInterview))

	http.HandleFunc("/interviews", interviewHandler)
	http.HandleFunc("/interviews/", interviewHandler)

	questionHandler := middleware.ValidateTokeMiddleware(
		middleware.ValidateUserMiddleware(
			controller.HandleQuestion))

	http.HandleFunc("/questions", questionHandler)
	http.HandleFunc("/questions/", questionHandler)
}

func main() {
	db.SetupDatabase()
	addRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
