package token

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"learning.go/appservice/config"
	"learning.go/appservice/model"
)

type ContextKey string

const JWT_KEY = ContextKey("JWT_CLAIM")

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func GetToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't retrive token")
	}
	return token, nil
}

func IsLoggedOut(jwtClaim model.JwtClaim) error {
	var URL = fmt.Sprintf("http://%s:8080/token/status/%s", config.SERVICE_NAME, jwtClaim.JTI)
	log.Println("calling..", URL)
	resp, err := http.Get(URL)
	if err != nil {
		return errors.Wrap(err, "Couldn't reach auth server")
	}
	if resp.StatusCode == http.StatusNotFound {
		return errors.Wrap(err, "User is logged out")
	}
	return nil
}
