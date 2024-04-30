package database

import (
	"fmt"
	cli_command "github.com/Xurliman/banking-microservice/internal/cli-command"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf("host=localhost user=%s dbname=%s password=%s port=5432 sslmode=disable", dbUser, dbName, dbPassword)

	db, err := gorm.Open(
		postgres.Open(connection), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	DB = db
	cli_command.Organize(db)
	return DB
}
