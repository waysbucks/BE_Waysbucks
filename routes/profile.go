package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product", h.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", h.UpdateProduct).Methods("PATCH")
	r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}
