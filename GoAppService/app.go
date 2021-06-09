package main

import (
	"log"
	"net/http"

	"learning.go/appservice/controller"
	"learning.go/appservice/db"
)

func addRouter() {
	http.HandleFunc("/openings", controller.HandleOpening)
	http.HandleFunc("/openings/", controller.HandleOpening)

	http.HandleFunc("/interviews", controller.HandleInterview)
	http.HandleFunc("/interviews/", controller.HandleInterview)

	http.HandleFunc("/questions", controller.HandleQuestion)
	http.HandleFunc("/questions/", controller.HandleQuestion)
}

func main() {
	db.SetupDatabase()
	addRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
