package controller

import (
	"net/http"

	"learning.go/authserver/model"
	"learning.go/authserver/token"
)

func HandleLogOut(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		jwtClaim := r.Context().Value("JWT_CLAIM").(model.JwtClaim)
		jti := jwtClaim.JTI
		token.GlobalTokenStore.Remove(jti)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
