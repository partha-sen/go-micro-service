package service

import (
	"fmt"

	"github.com/pkg/errors"
	"learning.go/appservice/dao"
	"learning.go/appservice/model"
)

func GetQuestionById(id int) (model.Any, error) {

	data, err := dao.GetQuestionById(id)
	if err != nil {
		return data, errors.Wrap(err, fmt.Sprintf("Couldn't retrieve Question for id  %d", id))
	}
	return data, nil
}

func GetAllQuestion() ([]model.Any, error) {

	data, err := dao.GetAllQuestions()

	if err != nil {
		return data, errors.Wrap(err, "Couldn't retrieve Question from database")
	}
	return data, nil

}

func SaveQuestion(obj model.Question) (int64, error) {

	id, err := dao.InsertQuestion(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't save record into database")
	}
	return id, nil
}

func UpdateQuestion(obj model.Question) (int64, error) {

	id, err := dao.UpdateQuestion(obj)

	if err != nil {
		return id, errors.Wrap(err, "Couldn't update record into database")
	}
	return id, nil
}

func DeleteQuestion(id int) (int64, error) {

	count, err := dao.DeleteQuestion(id)

	if err != nil {
		return count, errors.Wrap(err, "Couldn't delete record")
	}
	return count, nil
}
