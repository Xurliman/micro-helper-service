package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegistrationPlaceRepo struct {
	db *gorm.DB
}

func (repo *RegistrationPlaceRepo) Create(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error) {
	err := repo.db.Create(&registrationPlace).Error
	return registrationPlace, err
}

func (repo *RegistrationPlaceRepo) Get(id uuid.UUID) (models.RegistrationPlace, error) {
	var registrationPlace models.RegistrationPlace
	err := repo.db.Where("id = ?", id).First(&registrationPlace).Error

	return registrationPlace, err
}

func (repo *RegistrationPlaceRepo) Update(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error) {
	var dbRegistrationPlace models.RegistrationPlace
	if err := repo.db.Where("id = ?", registrationPlace.ID).First(&registrationPlace).Error; err != nil {
		return dbRegistrationPlace, err
	}

	dbRegistrationPlace.Code = registrationPlace.Code
	dbRegistrationPlace.Name = registrationPlace.Name
	dbRegistrationPlace.FlexFinId = registrationPlace.FlexFinId

	err := repo.db.Save(dbRegistrationPlace).Error
	return dbRegistrationPlace, err
}

func (repo *RegistrationPlaceRepo) GetByCode(code string) (models.RegistrationPlace, error) {
	var registrationPlace models.RegistrationPlace
	err := repo.db.Where("code = ?", code).First(&registrationPlace).Error

	return registrationPlace, err
}

func (repo *RegistrationPlaceRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.RegistrationPlace{}).Error
	return err
}
