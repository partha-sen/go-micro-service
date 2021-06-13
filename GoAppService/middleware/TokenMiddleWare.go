package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"learning.go/appservice/model"
	"learning.go/appservice/token"
)

func getJwtClaim(claims jwt.MapClaims) model.JwtClaim {

	var jwtClaim model.JwtClaim
	jwtClaim.UserName = claims["user_name"].(string)
	jwtClaim.Authorities = claims["authorities"].(string)
	jwtClaim.JTI = claims["jti"].(string)
	log.Println("jwtClaim ", jwtClaim)
	return jwtClaim
}

func ValidateAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Calling ValidateAdminMiddleware..")
		jwtClaim := r.Context().Value(token.JWT_KEY).(model.JwtClaim)
		if !strings.Contains(jwtClaim.Authorities, "ADMIN") {
			log.Println("Don't have ADMIN privilege")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next(w, r)
	})

}

func ValidateUserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Calling ValidateAdminMiddleware..")
		jwtClaim := r.Context().Value(token.JWT_KEY).(model.JwtClaim)
		if !strings.Contains(jwtClaim.Authorities, "USER") {
			log.Println("Don't have USER privilege")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next(w, r)
	})

}

func ValidateTokeMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Calling ValidateTokeMiddleware..")

		jwtToken, err := token.GetToken(r)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		claims, ok := jwtToken.Claims.(jwt.MapClaims)

		if !ok && !jwtToken.Valid {
			log.Println("invalid Token")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("claims ", claims)
		ctx := context.WithValue(r.Context(), token.JWT_KEY, getJwtClaim(claims))
		r = r.WithContext(ctx)
		next(w, r)
	})
}
