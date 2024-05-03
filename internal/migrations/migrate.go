package migrations

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
	"log"
)

type Migrator struct {
	db *gorm.DB
}

func Execute(connection *gorm.DB) {
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
