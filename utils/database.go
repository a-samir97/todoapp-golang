package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todoList/models"
)

func GetDatabaseConnection() *gorm.DB{
	db, err := gorm.Open("mysql", "ahmedsamir:root1234@/todolist?charset=utf8&parseTime=True&loc=Local")

	if err !=  nil {
		panic(err.Error())
	}
	//db.Debug().DropTableIfExists(&models.TodoItem{})
	db.Debug().AutoMigrate(&models.TodoItem{})
	return db
}
