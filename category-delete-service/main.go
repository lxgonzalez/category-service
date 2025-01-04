package main

import (
	"category-delete-service/database"
	"category-delete-service/handlers"
	"category-delete-service/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	r := mux.NewRouter()

	r.HandleFunc("/categories/{idCategory}", handlers.DeleteCategory).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
