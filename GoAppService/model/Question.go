package model

type Question struct {
	Id           int
	Interview_id int
	Tags         []string
	Text         string
}
