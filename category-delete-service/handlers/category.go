package handlers

import (
	"category-delete-service/database"
	"category-delete-service/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
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

	if err := publishToSNS(id); err != nil {
		log.Println("Error al publicar en SNS:", err)
		http.Error(w, "Failed to publish message to SNS: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func publishToSNS(categoryID string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Println("Error al crear la sesión de AWS:", err)
		return err
	}

	svc := sns.New(sess)

	message := map[string]string{
		"categoriaId": categoryID,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error al serializar el mensaje:", err)
		return err
	}

	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(messageBytes)),
		TopicArn: aws.String("arn:aws:sns:us-east-1:498431141888:categoryDelete"),
	})
	if err != nil {
		log.Println("Error al publicar en SNS:", err)
		return err
	}

	log.Println("Mensaje publicado en SNS:", *result.MessageId)
	return nil
}
