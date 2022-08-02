package main

import (
	"gorm-crud/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteUser).Methods("DELETE")
}
