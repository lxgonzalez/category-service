package main

import (
	"category-add-service/database"
	"category-add-service/handlers"
	"category-add-service/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	r := mux.NewRouter()

	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

	http.ListenAndServe(":1028", r)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Category Add Microservice is up and running ...")
}
