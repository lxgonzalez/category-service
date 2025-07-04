package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dbUser := os.Getenv("DATASOURCE_USERNAME")
	dbPassword := os.Getenv("DATASOURCE_PASSWORD")
	dbHost := os.Getenv("DATASOURCE_URL")
	dbPort := os.Getenv("DATASOURCE_PORT")
	dbName := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connection successful")
}
