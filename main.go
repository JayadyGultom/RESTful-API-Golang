package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"api/routes"
)

func main() {
	dsn := "root:@tcp(localhost:3306)/learn_"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to database")
	
	// Buat tabel students jika belum ada
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS students (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			age INT NOT NULL,
			grade VARCHAR(10) NOT NULL,
			address TEXT
		)
	`)
	if err != nil {
		panic(err.Error())
	}
	
	// Inisialisasi router
	router := mux.NewRouter()
	
	// Setup routes
	routes.SetupStudentRoutes(router, db)
	
	// Jalankan server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
