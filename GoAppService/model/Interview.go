package model

import "time"

type Interview struct {
	Id         int
	Opening_id int
	Date       time.Time
	Person     string
}
