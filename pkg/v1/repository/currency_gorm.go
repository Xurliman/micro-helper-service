package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
)

type CurrencyRepo struct {
	db *gorm.DB
}

func (repo CurrencyRepo) Create(currency models.Currency) (models.Currency, error) {
	err := repo.db.Create(&currency).Error
	return currency, err
}

func (repo CurrencyRepo) Get(id string) (models.Currency, error) {
	var currency models.Currency
	err := repo.db.Where("id = ?", id).First(&currency).Error

	return currency, err
}

func (repo CurrencyRepo) Update(currency models.Currency) (models.Currency, error) {
	var dbCurrency models.Currency
	if err := repo.db.Where("id = ?", currency.ID).First(&currency).Error; err != nil {
		return dbCurrency, err
	}

	dbCurrency.Code = currency.Code
	dbCurrency.Name = currency.Name
	dbCurrency.ShortName = currency.ShortName
	dbCurrency.Scale = currency.Scale
	dbCurrency.ScaleName = currency.ScaleName

	err := repo.db.Save(dbCurrency).Error
	return dbCurrency, err
}

func (repo CurrencyRepo) GetByCode(code string) (models.Currency, error) {
	var currency models.Currency
	err := repo.db.Where("code = ?", code).First(&currency).Error

	return currency, err
}

func (repo CurrencyRepo) Delete(id string) error {
	err := repo.db.Where("id = ?", id).Delete(&models.Currency{}).Error
	return err
}
