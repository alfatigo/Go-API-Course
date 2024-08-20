package routes

import (
	"encoding/json"
	"net/http"

	"github.com/alfatigo/Go-API-Course/db"
	"github.com/alfatigo/Go-API-Course/models"
)

func GetUsersHandler(w http.ResponseWriter, routes *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)

	// w.Write([]byte("Get Users"))
}

func GetUserHandler(w http.ResponseWriter, routes *http.Request) {
	w.Write([]byte("Get User"))
}
func PostUserHandler(w http.ResponseWriter, routes *http.Request) {
	var user models.User
	json.NewDecoder(routes.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(user)

	w.Write([]byte("Create User"))
}
func DeleteUsersHandler(w http.ResponseWriter, routes *http.Request) {
	w.Write([]byte("Delete User"))
}
