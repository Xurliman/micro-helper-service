package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NationalEconomySectorOldRepo struct {
	db *gorm.DB
}

func NewNationalEconomySectorOld(db *gorm.DB) interfaces.NationalEconomySectorOldRepoInterface {
	return &NationalEconomySectorOldRepo{db: db}
}

func (repo *NationalEconomySectorOldRepo) Create(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error) {
	err := repo.db.Create(&nationalEconomySectorOld).Error
	return nationalEconomySectorOld, err
}

func (repo *NationalEconomySectorOldRepo) Get(id uuid.UUID) (models.NationalEconomySectorOld, error) {
	var nationalEconomySectorOld models.NationalEconomySectorOld
	err := repo.db.Where("id = ?", id).First(&nationalEconomySectorOld).Error

	return nationalEconomySectorOld, err
}

func (repo *NationalEconomySectorOldRepo) Update(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error) {
	var dbNationalEconomySectorOld models.NationalEconomySectorOld
	if err := repo.db.Where("id = ?", nationalEconomySectorOld.ID).First(&nationalEconomySectorOld).Error; err != nil {
		return dbNationalEconomySectorOld, err
	}

	dbNationalEconomySectorOld.Code = nationalEconomySectorOld.Code
	dbNationalEconomySectorOld.Name = nationalEconomySectorOld.Name
	dbNationalEconomySectorOld.CBUCode = nationalEconomySectorOld.CBUCode
	dbNationalEconomySectorOld.CBUGroupCode = nationalEconomySectorOld.CBUGroupCode
	dbNationalEconomySectorOld.ActivationDate = nationalEconomySectorOld.ActivationDate
	dbNationalEconomySectorOld.DeactivationDate = nationalEconomySectorOld.DeactivationDate
	dbNationalEconomySectorOld.CBUReferenceKey = nationalEconomySectorOld.CBUReferenceKey
	dbNationalEconomySectorOld.FlexFinId = nationalEconomySectorOld.FlexFinId

	err := repo.db.Save(dbNationalEconomySectorOld).Error
	return dbNationalEconomySectorOld, err
}

func (repo *NationalEconomySectorOldRepo) GetByCode(code string) (models.NationalEconomySectorOld, error) {
	var nationalEconomySectorOld models.NationalEconomySectorOld
	err := repo.db.Where("code = ?", code).First(&nationalEconomySectorOld).Error

	return nationalEconomySectorOld, err
}

func (repo *NationalEconomySectorOldRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.NationalEconomySectorOld{}).Error
	return err
}
