package routes

import (
	"go-web-api/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}