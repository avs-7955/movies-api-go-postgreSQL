package controller

import (
	"database/sql"
	"fmt"
	"log"
	model "movies-api-go-post/models"
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

// postgreSQL helpers - file

// insert 1 record
func insertMovie(movie model.Netflix) {
	// fmt.Println(movie)
	ins_sql_statement := `INSERT INTO movies.netflix VALUES ($1, $2, $3) RETURNING id`
	id := 0
	// fmt.Println(db)
	err := db.QueryRow(ins_sql_statement, movie.Id, movie.Movie, movie.Watched).Scan(&id)
	CheckError(err)
	fmt.Println("\nRow inserted successfully!")
	fmt.Println("New record ID is:", id)
}

// update a movie to watched based on id
func updateMovie(id int) {
	upd_sql_statement := `UPDATE movies.netflix SET watched=true WHERE id=$1;`
	res, err := db.Exec(upd_sql_statement, id)
	CheckError(err)
	count, err := res.RowsAffected()

	CheckError(err)
	fmt.Printf("Rows updated: %v\n", count)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}