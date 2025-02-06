package main

import (
    "net/http"
    "category-list-service/database"
    "category-list-service/handlers"
    "category-list-service/models"
    "fmt"
    "github.com/gorilla/mux"
)

func main() {
    database.Connect()
    database.DB.AutoMigrate(&models.Category{})

    r := mux.NewRouter()

    r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
    r.HandleFunc("/categories/{idCategory}", handlers.GetOneCategory).Methods("GET")
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

    http.ListenAndServe(":1030", r)
}


func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Category List Microservice is up and running ...")
}