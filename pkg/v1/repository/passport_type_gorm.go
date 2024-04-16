package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PassportTypeRepo struct {
	db *gorm.DB
}

func (repo *PassportTypeRepo) Create(passportType models.PassportType) (models.PassportType, error) {
	err := repo.db.Create(&passportType).Error
	return passportType, err
}

func (repo *PassportTypeRepo) Get(id uuid.UUID) (models.PassportType, error) {
	var passportType models.PassportType
	err := repo.db.Where("id = ?", id).First(&passportType).Error

	return passportType, err
}

func (repo *PassportTypeRepo) Update(passportType models.PassportType) (models.PassportType, error) {
	var dbPassportType models.PassportType
	if err := repo.db.Where("id = ?", passportType.ID).First(&passportType).Error; err != nil {
		return dbPassportType, err
	}

	dbPassportType.Code = passportType.Code
	dbPassportType.Name = passportType.Name
	dbPassportType.OldCode = passportType.OldCode
	dbPassportType.OldName = passportType.OldName
	dbPassportType.ActivationDate = passportType.ActivationDate
	dbPassportType.DeactivationDate = passportType.DeactivationDate
	dbPassportType.FlexFinId = passportType.FlexFinId

	err := repo.db.Save(dbPassportType).Error
	return dbPassportType, err
}

func (repo *PassportTypeRepo) GetByCode(code string) (models.PassportType, error) {
	var passportType models.PassportType
	err := repo.db.Where("code = ?", code).First(&passportType).Error

	return passportType, err
}

func (repo *PassportTypeRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.PassportType{}).Error
	return err
}
