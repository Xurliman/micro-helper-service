package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
)

type NationalEconomySectorRepo struct {
	db *gorm.DB
}

func (repo *NationalEconomySectorRepo) Create(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error) {
	err := repo.db.Create(&nationalEconomySector).Error
	return nationalEconomySector, err
}

func (repo *NationalEconomySectorRepo) Get(id string) (models.NationalEconomySector, error) {
	var nationalEconomySector models.NationalEconomySector
	err := repo.db.Where("id = ?", id).First(&nationalEconomySector).Error

	return nationalEconomySector, err
}

func (repo *NationalEconomySectorRepo) Update(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error) {
	var dbNationalEconomySector models.NationalEconomySector
	if err := repo.db.Where("id = ?", nationalEconomySector.ID).First(&nationalEconomySector).Error; err != nil {
		return dbNationalEconomySector, err
	}

	dbNationalEconomySector.Code = nationalEconomySector.Code
	dbNationalEconomySector.Name = nationalEconomySector.Name
	dbNationalEconomySector.CBUCode = nationalEconomySector.CBUCode
	dbNationalEconomySector.CBUGroupCode = nationalEconomySector.CBUGroupCode
	dbNationalEconomySector.ActivationDate = nationalEconomySector.ActivationDate
	dbNationalEconomySector.DeactivationDate = nationalEconomySector.DeactivationDate
	dbNationalEconomySector.CBUReferenceKey = nationalEconomySector.CBUReferenceKey

	err := repo.db.Save(dbNationalEconomySector).Error
	return dbNationalEconomySector, err
}

func (repo *NationalEconomySectorRepo) GetByCode(code string) (models.NationalEconomySector, error) {
	var nationalEconomySector models.NationalEconomySector
	err := repo.db.Where("code = ?", code).First(&nationalEconomySector).Error

	return nationalEconomySector, err
}

func (repo *NationalEconomySectorRepo) Delete(id string) error {
	err := repo.db.Where("id = ?", id).Delete(&models.NationalEconomySector{}).Error
	return err
}
