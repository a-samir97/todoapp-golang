package models

type TodoItem struct {
	Id int `gorm:"primary_key"`
	Description string
	Completed bool
}
