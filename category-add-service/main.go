package main

import (
	"category-add-service/database"
	"category-add-service/handlers"
	"category-add-service/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	r := mux.NewRouter()

	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	http.ListenAndServe(":8080", r)
}
