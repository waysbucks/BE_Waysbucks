package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profiles", h.FindProfiles).Methods("GET")
	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
	r.HandleFunc("/profile", h.CreateProfile).Methods("POST")
	r.HandleFunc("/profile/{id}", h.UpdateProfile).Methods("PATCH")
	r.HandleFunc("/profile/{id}", h.DeleteProfile).Methods("DELETE")
}
