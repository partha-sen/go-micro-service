package controller

import (
	"net/http"

	"learning.go/authserver/token"
)

func HandleLogOut(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		jti := r.Header.Get("TOKEN_ID")
		token.GlobalTokenStore.Remove(jti)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
