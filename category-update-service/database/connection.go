package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	// Cargar variables desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtener las variables de entorno
	dbUser := os.Getenv("DATASOURCE_USERNAME")
	dbPassword := os.Getenv("DATASOURCE_PASSWORD")
	dbHost := os.Getenv("DATASOURCE_URL")
	dbPort := os.Getenv("DATASOURCE_PORT")
	dbName := os.Getenv("DATABASE")

	// Construir el Data Source Name (DSN) con las variables de entorno
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection successful")
}
