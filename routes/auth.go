package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authReposutory := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authReposutory)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}
