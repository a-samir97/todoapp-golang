package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"todoList/models"
	"todoList/utils"
)
var db = utils.GetDatabaseConnection()

func APIHealth(w http.ResponseWriter, r *http.Request) {
	log.Info("API is Ok")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello World"))
	io.WriteString(w, `{"alive": true}`)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	todo := &models.TodoItem{Description: description, Completed: false}
	db.Create(&todo)
	result := db.Last(&todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Get Id from URL Parameter
	vars := mux.Vars(r)
	// to convert from string to int
	id, _ := strconv.Atoi(vars["id"])

	getId := utils.GetItemByID(id, db)

	if getId {
		completed, _ := strconv.ParseBool(r.FormValue("completed"))
		log.WithFields(log.Fields{"Id": id, "Completed": completed}).Info("Updating TodoItem")
		todo := &models.TodoItem{}
		db.First(&todo, id)
		todo.Completed = completed
		db.Save(&todo)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Updated Successfully!"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id does not exist in our database"))
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	getId := utils.GetItemByID(id, db)

	if getId {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting Item")
		todo := &models.TodoItem{}
		db.First(&todo, id)
		db.Delete(&todo)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Deleted Successfully!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Id does not exist in our database"))
	}
}

func GetCompletedItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Completed Items")
	completedTodoItems := utils.GetTodoItemsByCompleted(true, db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completedTodoItems)
}

func GetIncompletedItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Get incompleted items")
	incomepletedTodoItems := utils.GetTodoItemsByCompleted(false, db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomepletedTodoItems)
}