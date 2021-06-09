package dao

import (
	"learning.go/appservice/db"
	"learning.go/appservice/model"
)

func InsertQuestion(q model.Question) (int64, error) {
	conn := db.GetConnection()
	result, err := conn.Exec(`INSERT INTO question 
	(interview_id, 
		text) 
		VALUES (?, ?)`,
		q.Interview_id,
		q.Text)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetQuestionById(id int) (model.Any, error) {

	conn := db.GetConnection()
	var q model.Question

	row := conn.QueryRow(` SELECT id, 
	interview_id, 
	text
	FROM question 
	WHERE id = ?`, id)

	err := row.Scan(&q.Id, &q.Interview_id, &q.Text)
	if err != nil {
		return q, err
	}

	return q, nil
}

func GetAllQuestions() ([]model.Any, error) {

	allQuestions := []model.Any{}

	conn := db.GetConnection()
	results, err := conn.Query(`SELECT id, 
	interview_id, 
	text
	FROM question`)

	if err != nil {
		return allQuestions, err
	}

	defer results.Close()

	for results.Next() {
		var q model.Question
		results.Scan(&q.Id, &q.Interview_id, &q.Text)
		allQuestions = append(allQuestions, q)
	}

	return allQuestions, nil

}

func UpdateQuestion(q model.Question) (int64, error) {
	conn := db.GetConnection()
	result, err := conn.Exec(`UPDATE question 
	  set text=? 
		WHERE id=?`,
		q.Text,
		q.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func DeleteQuestion(id int) (int64, error) {

	conn := db.GetConnection()
	result, err := conn.Exec(`DELETE FROM question WHERE id=?`, id)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}
