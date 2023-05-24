package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	model "movies-api-go-post/models"
	"net/http"
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
	// returns id of the record being created
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

// delete a movie based on id
func deleteMovie(id int) {
	del_sql_statement := `DELETE FROM movies.netflix WHERE id=$1;`
	res, err := db.Exec(del_sql_statement, id)
	CheckError(err)
	count, err := res.RowsAffected()

	CheckError(err)
	fmt.Printf("Rows deleted: %v\n", count)
}

// delete all movies
func deleteMovies() int {
	del_sql_statement := `DELETE FROM movies.netflix;`
	res, err := db.Exec(del_sql_statement)
	CheckError(err)
	count, err := res.RowsAffected()

	CheckError(err)
	fmt.Printf("No. of movie records deleted: %v\n", count)
	return int(count)
}

// retriving all movies from the database
func getMovies() []model.Netflix {
	rtv_sql_statement := `SELECT * FROM movies.netflix;`
	rows, err := db.Query(rtv_sql_statement)
	CheckError(err)
	defer rows.Close()

	// creates placeholder of the returned records
	movie_records := make([]model.Netflix, 0)

	// storing the rows into structures
	for rows.Next() {
		record := model.Netflix{}
		err = rows.Scan(&record.Id, &record.Movie, &record.Watched)
		CheckError(err)
		movie_records = append(movie_records, record)
	}
	return movie_records
}

// actual controllers - file
func GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all movies")
	w.Header().Set("Content-Type", "application/json") // setting headers to the content
	allmovies := getMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Insert a movie")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") // will only accept post requests
	// if the data received is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie) // decoding the data received
	CheckError(err)
	// rand.Seed(time.Now().UnixNano())
	movie.Id = rand.Intn(500) // adding a primary key
	fmt.Println(movie)

	insertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
