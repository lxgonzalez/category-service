package main

import (
	"category-update-service/database"
	"category-update-service/handlers"
	"category-update-service/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Category{})

	r := mux.NewRouter()

	r.HandleFunc("/categories", handlers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/", HealthCheckHandler).Methods("GET")

	http.ListenAndServe(":1031", r)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Category Update Microservice is up and running ...")
}
