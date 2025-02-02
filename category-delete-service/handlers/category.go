package handlers

import (
	"category-delete-service/database"
	"category-delete-service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func notifyBroker(categoryID string) {
	brokerURL := os.Getenv("BROKER_URL")

	conn, _, err := websocket.DefaultDialer.Dial(brokerURL, nil)
	if err != nil {
		fmt.Println("Error connecting to broker:", err)
		return
	}
	defer conn.Close()

	message := fmt.Sprintf(`{"event": "delete_category", "category_id": "%s"}`, categoryID)
	fmt.Println("Sending message to broker:", message)
	conn.WriteMessage(websocket.TextMessage, []byte(message))
}

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

	notifyBroker(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
