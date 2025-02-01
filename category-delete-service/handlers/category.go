package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Crea una sesión de AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Cambia a tu región
	})
	if err != nil {
		log.Fatal("Error al crear la sesión de AWS:", err)
	}

	// Crea un cliente de SNS
	svc := sns.New(sess)

	// Define el mensaje
	message := map[string]string{
		"categoriaId": "123", // Cambia por el ID de la categoría
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Error al serializar el mensaje:", err)
	}

	// Publica el mensaje en el Topic de SNS
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(messageBytes)),
		TopicArn: aws.String("arn:aws:sns:us-east-1:498431141888:categoryDelete"), // Cambia por tu ARN
	})
	if err != nil {
		log.Fatal("Error al publicar en SNS:", err)
	}

	log.Println("Mensaje publicado en SNS:", *result.MessageId)
}
