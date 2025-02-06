package main

import (
	"category-delete-service/database"
	"category-delete-service/handlers"
	"category-delete-service/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	r := mux.NewRouter()

	r.HandleFunc("/categories/{idCategory}", handlers.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

	http.ListenAndServe(":1029", r)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Category Delete Microservice is up and running ...")
}
