package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
)

type BankRepo struct {
	db *gorm.DB
}

func (repo *BankRepo) Create(bank models.Bank) (models.Bank, error) {
	err := repo.db.Create(&bank).Error
	return bank, err
}

func (repo *BankRepo) Get(id string) (models.Bank, error) {
	var bank models.Bank
	err := repo.db.Where("id = ?", id).First(&bank).Error

	return bank, err
}

func (repo *BankRepo) Update(bank models.Bank) (models.Bank, error) {
	var dbBank models.Bank
	if err := repo.db.Where("id = ?", bank.ID).First(&bank).Error; err != nil {
		return dbBank, err
	}

	dbBank.Code = bank.Code
	dbBank.Name = bank.Name
	dbBank.ShortName = bank.ShortName
	dbBank.CountryID = bank.CountryID
	dbBank.OpenDate = bank.OpenDate
	dbBank.CloseDate = bank.CloseDate
	dbBank.ActivationDate = bank.ActivationDate
	dbBank.DeactivationDate = bank.DeactivationDate

	err := repo.db.Save(dbBank).Error
	return dbBank, err
}

func (repo *BankRepo) Delete(id string) error {
	err := repo.db.Where("id = ?", id).Delete(&models.Bank{}).Error
	return err
}

func (repo *BankRepo) GetByCode(code string) (models.Bank, error) {
	var bank models.Bank
	err := repo.db.Where("code = ?", code).First(&bank).Error

	return bank, err
}

func NewBank(db *gorm.DB) interfaces.BankRepoInterface {
	return &BankRepo{db: db}
}
