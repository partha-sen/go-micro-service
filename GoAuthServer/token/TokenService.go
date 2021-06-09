package token

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"learning.go/authserver/config"
)

func CreateToken(userid string, authorities []string) (string, error) {
	var err error
	var expiryTime int64
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = userid
	log.Println("authorities ", authorities)
	if authoritiesStr, err := json.Marshal(authorities); err == nil {
		atClaims["authorities"] = string(authoritiesStr)
	}
	expiryTime = time.Now().Add(time.Minute * config.TOKEN_VALIDITY).Unix()
	atClaims["exp"] = expiryTime
	jtiUuid := uuid.New().String()
	atClaims["jti"] = jtiUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	generatedToken, err := at.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return "", err
	}
	log.Println("JTI", jtiUuid)
	GlobalTokenStore.Add(jtiUuid, expiryTime)
	return generatedToken, nil
}

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
