package main

import (
	"fmt"
	"log"
	"movies-api-go-post/router"
	"net/http"

	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("PostgreSQL Movies API")
	r := router.Router()
	fmt.Println("Server is getting started....")
	// listening to a port
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening at port 3000....")
}
