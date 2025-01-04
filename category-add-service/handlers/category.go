package handlers

import (
    "encoding/json"
    "net/http"
    "category-add-service/database"
    "category-add-service/models"

)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    database.DB.Create(&category)
    json.NewEncoder(w).Encode(category)
}
