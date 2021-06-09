package service

import (
	"fmt"

	"github.com/pkg/errors"
	"learning.go/appservice/dao"
	"learning.go/appservice/model"
)

func GetInterviewById(id int) (model.Any, error) {

	data, err := dao.GetInterviewById(id)
	if err != nil {
		return data, errors.Wrap(err, fmt.Sprintf("Couldn't retrieve Interview for id  %d", id))
	}
	return data, nil
}

func GetAllInterview() ([]model.Any, error) {

	data, err := dao.GetAllInterviews()

	if err != nil {
		return data, errors.Wrap(err, "Couldn't retrieve Interview from database")
	}
	return data, nil

}

func SaveInterview(obj model.Interview) (int64, error) {

	id, err := dao.InsertInterview(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't save record into database")
	}
	return id, nil
}

func UpdateInterview(obj model.Interview) (int64, error) {

	id, err := dao.UpdateInterview(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't update record into database")
	}
	return id, nil
}

func DeleteInterview(id int) (int64, error) {

	count, err := dao.DeleteInterview(id)

	if err != nil {
		return count, errors.Wrap(err, "Couldn't delete record")
	}
	return count, nil
}
