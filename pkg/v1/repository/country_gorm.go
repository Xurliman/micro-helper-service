package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CountryRepo struct {
	db *gorm.DB
}

func NewCountry(db *gorm.DB) interfaces.CountryRepoInterface {
	return &CountryRepo{db: db}
}

func (repo *CountryRepo) Create(country models.Country) (models.Country, error) {
	err := repo.db.Create(&country).Error
	return country, err
}

func (repo *CountryRepo) Get(id uuid.UUID) (models.Country, error) {
	var country models.Country
	err := repo.db.Where("id = ?", id).First(&country).Error

	return country, err
}

func (repo *CountryRepo) Update(country models.Country) (models.Country, error) {
	var dbCountry models.Country
	if err := repo.db.Where("id = ?", country.ID).First(&country).Error; err != nil {
		return dbCountry, err
	}

	dbCountry.Code = country.Code
	dbCountry.Name = country.Name
	dbCountry.ShortName = country.ShortName
	dbCountry.CurrencyId = country.CurrencyId
	dbCountry.CodeAlpha2 = country.CodeAlpha2
	dbCountry.CodeAlpha3 = country.CodeAlpha3
	dbCountry.TerritoryCode = country.TerritoryCode
	dbCountry.ActivationDate = country.ActivationDate
	dbCountry.DeactivationDate = country.DeactivationDate
	dbCountry.FlexFinId = country.FlexFinId

	err := repo.db.Save(dbCountry).Error
	return dbCountry, err
}

func (repo *CountryRepo) GetByCode(code string) (models.Country, error) {
	var country models.Country
	err := repo.db.Where("code = ?", code).First(&country).Error

	return country, err
}

func (repo *CountryRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.Country{}).Error
	return err
}
