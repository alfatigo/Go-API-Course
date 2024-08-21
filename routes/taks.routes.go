package routes

import (
	"encoding/json"
	"net/http"

	"github.com/alfatigo/Go-API-Course/db"
	"github.com/alfatigo/Go-API-Course/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Tasks
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)

}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Taks not Found!"))
		return
	}
	json.NewEncoder(w).Encode(&task)

}
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	json.NewDecoder(r.Body).Decode(&task)

	createTask := db.DB.Create(&task)
	err := createTask.Error
	if err != nil {
		w.WriteHeader((http.StatusBadRequest))
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var taks models.User
	params := mux.Vars(r)
	db.DB.First(&taks, params["id"])
	if taks.ID == 0 {

		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not Found!"))
		return
	}
	db.DB.Unscoped().Delete(&taks)
	w.WriteHeader(http.StatusOK)
}
