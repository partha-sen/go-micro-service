package dao

import (
	"log"

	"learning.go/appservice/db"
	"learning.go/appservice/model"
)

func InsertInterview(it model.Interview) (int64, error) {
	conn := db.GetConnection()
	result, err := conn.Exec(`INSERT INTO Interview 
	(opening_id, 
		date, 
		person) 
		VALUES (?, ?, ?)`,
		it.Opening_id,
		it.Date,
		it.Person)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetInterviewById(id int) (model.Any, error) {

	conn := db.GetConnection()
	it := model.Interview{}

	row := conn.QueryRow(` SELECT id, 
	opening_id, 
	date, 
	person
	FROM Interview 
	WHERE id = ?`, id)

	err := row.Scan(&it.Id, &it.Opening_id, &it.Date, &it.Person)
	if err != nil {
		return it, err
	}

	return it, nil

}

func GetAllInterviews() ([]model.Any, error) {

	allInterviews := []model.Any{}

	conn := db.GetConnection()
	results, err := conn.Query(`SELECT id, 
	opening_id, 
	date, 
	person
	FROM Interview`)

	if err != nil {
		return allInterviews, err
	}

	defer results.Close()

	for results.Next() {
		var it model.Interview
		results.Scan(&it.Id, &it.Opening_id, &it.Date, &it.Person)
		allInterviews = append(allInterviews, it)
	}

	return allInterviews, nil

}

func UpdateInterview(it model.Interview) (int64, error) {
	conn := db.GetConnection()
	result, err := conn.Exec(`UPDATE Interview 
	  set date=?, 
	    person=? 
		WHERE id=?`,
		it.Date,
		it.Person,
		it.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func DeleteInterview(id int) (int64, error) {

	log.Printf("DELETE FROM Interview WHERE id=%v \n", id)

	conn := db.GetConnection()
	result, err := conn.Exec(`DELETE FROM Interview WHERE id=?`, id)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}
