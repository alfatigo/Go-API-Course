package main

import (
	"net/http"

	"github.com/alfatigo/Go-API-Course/db"
	"github.com/alfatigo/Go-API-Course/models"
	"github.com/alfatigo/Go-API-Course/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Tasks{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	//Users Roter
	r.HandleFunc("/", routes.Homehandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//Tasks Router
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.AddTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
