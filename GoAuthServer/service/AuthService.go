package service

import (
	"github.com/pkg/errors"
	"learning.go/authserver/dao"
	"learning.go/authserver/model"
	"learning.go/authserver/token"
)

func IssueToken(user model.User) (string, error) {
	if usr, err := dao.ValidateAndGetUser(user); err != nil {
		return "", errors.Wrap(err, "Can't issue token")
	} else {
		token, err := token.CreateToken(usr.Username, usr.Role)
		return token, err
	}
}
