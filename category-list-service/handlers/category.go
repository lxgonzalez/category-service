package handlers

import (
    "encoding/json"
    "net/http"
    "category-list-service/database"
    "category-list-service/models"
    "github.com/gorilla/mux"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
    var categories []models.Category
    database.DB.Find(&categories)
    json.NewEncoder(w).Encode(categories)
}

func GetOneCategory(w http.ResponseWriter, r *http.Request) {
    // Obtém o ID da categoria dos parâmetros da URL
    vars := mux.Vars(r)
    id := vars["idCategory"]

    var category models.Category
    if err := database.DB.First(&category, id).Error; err != nil {
        http.Error(w, "Category not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(category)
}