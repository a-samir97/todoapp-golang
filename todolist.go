package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"todoList/handlers"
	"todoList/utils"
)

func init() {
	fmt.Println("Hello from init function")
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	db := utils.GetDatabaseConnection()
	defer db.Close()

	log.Info("Starting API Server at port 8000")

	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.APIHealth).Methods("GET")
	router.HandleFunc("/todo/completed", handlers.GetCompletedItems).Methods("GET")
	router.HandleFunc("/todo/incompleted", handlers.GetIncompletedItems).Methods("GET")

	router.HandleFunc("/todo", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/todo/{id}", handlers.DeleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}