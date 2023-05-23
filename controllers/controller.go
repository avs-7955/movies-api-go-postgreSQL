package controller

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

var db *sql.DB // pointer to the database

func init() {
	// establishing connection to the database
	fmt.Println("Connecting to db.... ")
	// connection string
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)

	// open database
	var err error
	db, err = sql.Open("postgres", connStr)
	CheckError(err)
	// defer db.Close()

	//check database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nSuccessfully connected to the database!")

	// creating movies schema
	command := `CREATE SCHEMA movies ;`
	_, err = db.Exec(command)
	CheckError(err)
	fmt.Println("\nMovies schema created.")

	// creating table inside schema
	command = `CREATE TABLE movies.netflix(id INT PRIMARY KEY, movie_name VARCHAR(240),watched BOOLEAN);`
	_, err = db.Exec(command)
	CheckError(err)
	fmt.Println("Movies table created successfully.")
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
