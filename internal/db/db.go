package database

import (
	"fmt"
	"github.com/Xurliman/banking-microservice/internal/models"
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
	AutoMigrate(db)
	return DB
}

func AutoMigrate(connection *gorm.DB) {
	err := connection.AutoMigrate(
		&models.Account{},
		&models.Bank{},
		&models.BankBranch{},
		&models.BorrowerType{},
		&models.ClientTypeClassifier{},
		&models.Country{},
		&models.Currency{},
		&models.DirectOrgan{},
		&models.District{},
		&models.EducationType{},
		&models.NationalEconomySectorNew{},
		&models.NationalEconomySectorOld{},
		&models.PassportType{},
		&models.PaymentType{},
		&models.Region{},
		&models.RegistrationPlace{},
		&models.ResidencyType{},
		&models.TaxOrganisation{},
	)
	if err != nil {
		log.Printf("err when automigrate: %v", err)
	}
}
