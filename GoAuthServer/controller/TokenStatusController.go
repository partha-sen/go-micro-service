package controller

import (
	"log"
	"net/http"
	"strings"

	"learning.go/authserver/token"
)

func HandleTokenValidity(w http.ResponseWriter, r *http.Request) {

	strId := strings.TrimPrefix(r.URL.Path, "/token/status/")

	log.Println("Token Id", strId)

	switch r.Method {

	case http.MethodGet:
		if token.GlobalTokenStore.Contains(strId) {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
