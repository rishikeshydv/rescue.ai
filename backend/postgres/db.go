package postgres

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func DBConnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//getting credentials for postgresql
	server := os.Getenv("POSTGRES_SERVER")
	database := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	log.Println(server, database, port, username, password)

}
