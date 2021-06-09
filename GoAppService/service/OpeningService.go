package service

import (
	"fmt"

	"github.com/pkg/errors"
	"learning.go/appservice/dao"
	"learning.go/appservice/model"
)

func GetOpeningById(id int) (model.Any, error) {

	data, err := dao.GetOpeningById(id)
	if err != nil {
		return data, errors.Wrap(err, fmt.Sprintf("Couldn't retrieve Opening for id  %d", id))
	}
	return data, nil
}

func GetAllOpening() ([]model.Any, error) {

	data, err := dao.GetAllOpening()

	if err != nil {
		return data, errors.Wrap(err, "Couldn't retrieve Opening from database")
	}
	return data, nil

}

func SaveOpening(obj model.Opening) (int64, error) {

	id, err := dao.InsertOpening(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't save record into database")
	}
	return id, nil
}

func UpdateOpening(obj model.Opening) (int64, error) {

	id, err := dao.UpdateOpening(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't update record into database")
	}
	return id, nil
}

func DeleteOpening(id int) (int64, error) {

	count, err := dao.DeleteOpening(id)

	if err != nil {
		return count, errors.Wrap(err, "Couldn't delete record")
	}
	return count, nil
}
