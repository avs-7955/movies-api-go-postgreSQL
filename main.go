package main

import (
	"fmt"

	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("PostgreSQL Movies API")
}
