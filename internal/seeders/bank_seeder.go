package seeds

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
)

func (seed Seed) BankSeeder() {
	transaction := seed.db.Begin()
	defer func() {
		if recovered := recover(); recovered != nil {
			log.Printf("There is an error so we are rolling back")
			transaction.Rollback()
		}
	}()

	if transaction.Error != nil {
		log.Fatalf("There is a transaction error %v", transaction.Error)
	}

	var banks []*models.Bank
	for i := 0; i < 3; i++ {
		newUUID, _ := uuid.NewUUID()
		banks = append(banks, &models.Bank{
			Model: models.Model{
				ID:        newUUID,
				CreatedAt: faker.Date(),
				UpdatedAt: faker.Date(),
				DeletedAt: gorm.DeletedAt{},
			},
			Code:             faker.Word(),
			Name:             faker.Name(),
			ShortName:        faker.LastName(),
			CountryID:        int64(rand.Int()),
			OpenDate:         faker.Date(),
			CloseDate:        faker.Date(),
			ActivationDate:   faker.Date(),
			DeactivationDate: faker.Date(),
			FlexFinId:        strconv.Itoa(rand.Int()),
		})
	}
	for _, bank := range banks {
		log.Printf("bank %v", bank)
	}
	seed.db.Create(banks)
	transaction.Commit()
}
