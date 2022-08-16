package main

import (
	"fmt"
	"net/http"
	"waysbucks/database"
	"waysbucks/pkg/mysql"
	"waysbucks/routes"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	//try
	fmt.Println("Server Running on localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
