package controller

import (
	"log"
	"net/http"

	"learning.go/authserver/model"
	"learning.go/authserver/service"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		obj := model.User{}
		requestBodyToObject(w, r, &obj)
		token, err := service.IssueToken(obj)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusForbidden)
		}
		w.Write([]byte(token))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
