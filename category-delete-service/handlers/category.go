package handlers

import (
    "encoding/json"
    "net/http"
    "category-delete-service/database"
    "category-delete-service/models"
    "github.com/gorilla/mux"

)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["idCategory"]

    var category models.Category
    if err := database.DB.First(&category, id).Error; err != nil {
        http.Error(w, "Category not found", http.StatusNotFound)
        return
    }

    if err := database.DB.Delete(&category).Error; err != nil {
        http.Error(w, "Failed to delete category", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(category)
}