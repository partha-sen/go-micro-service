package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"learning.go/authserver/token"
)

func setHedder(r *http.Request, claims jwt.MapClaims) {
	r.Header.Add("USER_NAMR", claims["user_name"].(string))
	r.Header.Add("USER_ROLE", claims["authorities"].(string))
	r.Header.Add("TOKEN_ID", claims["jti"].(string))
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
			setHedder(r, claims)
			next(w, r)
		}
	})
}
