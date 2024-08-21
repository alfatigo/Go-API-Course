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

	route := mux.NewRouter()
	//Routes
	route.HandleFunc("/", routes.Homehandler)
	route.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	route.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	route.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	route.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", route)
}
