package database

import (
	"fmt"
	"github.com/Xurliman/banking-microservice/config"
	cli_command "github.com/Xurliman/banking-microservice/internal/cli-command"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DbConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	connection := fmt.Sprintf("host=localhost user=%s dbname=%s password=%s port=5432 sslmode=disable", config.GetDatabaseUser(), config.GetDatabaseName(), config.GetDatabasePassword())

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
