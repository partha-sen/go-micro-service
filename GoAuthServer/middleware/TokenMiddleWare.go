package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"learning.go/authserver/model"
	"learning.go/authserver/token"
)

func getJwtClaim(claims jwt.MapClaims) model.JwtClaim {

	var jwtClaim model.JwtClaim
	jwtClaim.UserName = claims["user_name"].(string)
	jwtClaim.Authorities = claims["authorities"].(string)
	jwtClaim.JTI = claims["jti"].(string)
	return jwtClaim
}

func ValidateTokeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token, err := token.GetToken(r); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok && !token.Valid {
				log.Println("Couldn't retrieve role from Token")
				w.WriteHeader(http.StatusBadRequest)
			}
			ctx := context.WithValue(r.Context(), "JWT_CLAIM", getJwtClaim(claims))
			r = r.WithContext(ctx)
			next(w, r)
		}
	})
}
