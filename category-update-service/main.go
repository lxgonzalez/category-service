package main

import (
    "net/http"
    "category-update-service/database"
    "category-update-service/handlers"
    "category-update-service/models"

    "github.com/gorilla/mux"
)

func main() {
    database.Connect()
    database.DB.AutoMigrate(&models.Category{})

    r := mux.NewRouter()

    r.HandleFunc("/categories", handlers.UpdateCategory).Methods("PUT")

    http.ListenAndServe(":8080", r)
}
