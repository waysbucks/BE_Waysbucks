package main

import (
	"fmt"
	"net/http"
	"waysbucks/database"
	"waysbucks/pkg/mysql"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	fmt.Println("Server Running on localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
