package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", h.FindToppings).Methods("GET")
	r.HandleFunc("/topping/{id}", h.GetTopping).Methods("GET")
	r.HandleFunc("/topping", h.CreateTopping).Methods("POST")
	r.HandleFunc("/topping/{id}", h.UpdateTopping).Methods("PATCH")
	r.HandleFunc("/topping/{id}", h.DeleteTopping).Methods("DELETE")
}
