package dao

import (
	"learning.go/appservice/db"
	"learning.go/appservice/model"
)

func InsertOpening(o model.Opening) (int64, error) {
	conn := db.GetConnection()
	result, err := conn.Exec(`INSERT INTO Opening 
	(name, 
		company, 
		position, 
		skills, 
		experience) 
		VALUES (?, ?, ?, ?, ?)`,
		o.Name,
		o.Company,
		o.Position,
		o.Skills,
		o.Experience)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetOpeningById(id int) (model.Any, error) {

	conn := db.GetConnection()
	op := model.Opening{}

	row := conn.QueryRow(` SELECT id, 
	name, 
	company, 
	POSITION, 
	skills, 
	experience 
	FROM Opening 
	WHERE id = ?`, id)

	err := row.Scan(&op.Id, &op.Name, &op.Company, &op.Position, &op.Skills, &op.Experience)
	if err != nil {
		return op, err
	}

	return op, nil

}

func GetAllOpening() ([]model.Any, error) {

	allOpening := []model.Any{}

	conn := db.GetConnection()
	results, err := conn.Query("SELECT id, NAME, company, POSITION, skills, experience FROM Opening")

	if err != nil {
		return allOpening, err
	}

	defer results.Close()

	for results.Next() {
		var op model.Opening
		results.Scan(&op.Id, &op.Name, &op.Company, &op.Position, &op.Skills, &op.Experience)
		allOpening = append(allOpening, op)
	}

	return allOpening, nil

}

func UpdateOpening(op model.Opening) (int64, error) {
	conn := db.GetConnection()
	result, err := conn.Exec(`UPDATE Opening 
	  set name=?, 
		company=?, 
		position=?, 
		skills=?, 
		experience=?
		WHERE id=?`,
		op.Name,
		op.Company,
		op.Position,
		op.Skills,
		op.Experience,
		op.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func DeleteOpening(id int) (int64, error) {

	conn := db.GetConnection()
	result, err := conn.Exec(`DELETE FROM Opening WHERE  id=?`, id)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}
