package main

import (
    "net/http"
    "category-list-service/database"
    "category-list-service/handlers"
    "category-list-service/models"

    "github.com/gorilla/mux"
)

func main() {
    database.Connect()
    database.DB.AutoMigrate(&models.Category{})

    r := mux.NewRouter()

    r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
    r.HandleFunc("/categories/{idCategory}", handlers.GetOneCategory).Methods("GET")

    http.ListenAndServe(":8080", r)
}
