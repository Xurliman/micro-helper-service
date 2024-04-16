package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResidencyTypeRepo struct {
	db *gorm.DB
}

func (repo *ResidencyTypeRepo) Create(residencyType models.ResidencyType) (models.ResidencyType, error) {
	err := repo.db.Create(&residencyType).Error
	return residencyType, err
}

func (repo *ResidencyTypeRepo) Get(id uuid.UUID) (models.ResidencyType, error) {
	var residencyType models.ResidencyType
	err := repo.db.Where("id = ?", id).First(&residencyType).Error

	return residencyType, err
}

func (repo *ResidencyTypeRepo) Update(residencyType models.ResidencyType) (models.ResidencyType, error) {
	var dbResidencyType models.ResidencyType
	if err := repo.db.Where("id = ?", residencyType.ID).First(&residencyType).Error; err != nil {
		return dbResidencyType, err
	}

	dbResidencyType.Code = residencyType.Code
	dbResidencyType.Name = residencyType.Name
	dbResidencyType.ActivationDate = residencyType.ActivationDate
	dbResidencyType.DeactivationDate = residencyType.DeactivationDate
	dbResidencyType.FlexFinId = residencyType.FlexFinId
	dbResidencyType.NameUz = residencyType.NameUz

	err := repo.db.Save(dbResidencyType).Error
	return dbResidencyType, err
}

func (repo *ResidencyTypeRepo) GetByCode(code string) (models.ResidencyType, error) {
	var residencyType models.ResidencyType
	err := repo.db.Where("code = ?", code).First(&residencyType).Error

	return residencyType, err
}

func (repo *ResidencyTypeRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.ResidencyType{}).Error
	return err
}
