package dao

import (
	"errors"

	"learning.go/authserver/db"
	"learning.go/authserver/model"
)

func ValidateAndGetUser(usr model.User) (model.User, error) {
	if user, err := getUserId(usr); err == nil {
		if user.Password == usr.Password {
			return user, nil
		} else {
			return user, errors.New("Password not match")
		}
	} else {
		return user, errors.New("User not found")
	}
}

func getUserId(usr model.User) (model.User, error) {

	conn := db.GetConnection()
	entity := model.User{}

	row := conn.QueryRow(`SELECT user_name, 
	password
	FROM user 
	WHERE user_name =?`, usr.Username)

	err := row.Scan(&entity.Username, &entity.Password)
	if err != nil {
		return entity, err
	}

	results, err := conn.Query(`SELECT role_name FROM user_role, role
		WHERE user_role.role_id=role.id
		AND user_role.user_name=?`, usr.Username)
	defer results.Close()

	if err == nil {
		roles := []string{}

		for i := 0; results.Next(); i++ {
			var role string
			results.Scan(&role)
			roles = append(roles, role)
		}

		entity.Role = roles
	}

	return entity, nil

}
