package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"api/controllers"
)

func SetupStudentRoutes(router *mux.Router, db *sql.DB) {
	studentController := &controllers.StudentController{DB: db}
	
	// Definisikan endpoint untuk operasi CRUD
	router.HandleFunc("/students", studentController.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", studentController.GetStudent).Methods("GET")
	router.HandleFunc("/students", studentController.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id}", studentController.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", studentController.DeleteStudent).Methods("DELETE")
}
