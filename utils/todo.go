package utils

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"todoList/models"
)

func GetItemByID(Id int, db *gorm.DB) bool {
	todo := &models.TodoItem{}
	result := db.First(&todo, Id)
	if result.Error != nil {
		log.Warn("Todo item not found in database")
		return false
	}
	return true
}

func GetTodoItemsByCompleted(completed bool, db *gorm.DB) interface{} {
	var todos [] models.TodoItem
	TodoItems := db.Where("completed = ?", completed).Find(&todos).Value
	return TodoItems
}