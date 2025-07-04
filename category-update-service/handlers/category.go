package handlers

import (
    "encoding/json"
    "net/http"
    "category-update-service/database"
    "category-update-service/models"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
    var updatedData models.Category

    if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    var existingCategory models.Category
    if err := database.DB.First(&existingCategory, updatedData.IDCategory).Error; err != nil {
        http.Error(w, "Category not found", http.StatusNotFound)
        return
    }

    existingCategory.Name = updatedData.Name

    if err := database.DB.Save(&existingCategory).Error; err != nil {
        http.Error(w, "Failed to update category", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(existingCategory)
}