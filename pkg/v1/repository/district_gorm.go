package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
)

type DistrictRepo struct {
	db *gorm.DB
}

func (repo *DistrictRepo) Create(district models.District) (models.District, error) {
	err := repo.db.Create(&district).Error
	return district, err
}

func (repo *DistrictRepo) Get(id int64) (models.District, error) {
	var district models.District
	err := repo.db.Where("id = ?", id).First(&district).Error

	return district, err
}

func (repo *DistrictRepo) Update(district models.District) (models.District, error) {
	var dbDistrict models.District
	if err := repo.db.Where("id = ?", district.ID).First(&district).Error; err != nil {
		return dbDistrict, err
	}

	dbDistrict.Code = district.Code
	dbDistrict.Name = district.Name
	dbDistrict.RegionID = district.RegionID
	dbDistrict.ActivationDate = district.ActivationDate
	dbDistrict.DeactivationDate = district.DeactivationDate
	dbDistrict.OldCode = district.OldCode

	err := repo.db.Save(dbDistrict).Error
	return dbDistrict, err
}

func (repo *DistrictRepo) GetByCode(code string) (models.District, error) {
	var district models.District
	err := repo.db.Where("code = ?", code).First(&district).Error

	return district, err
}

func (repo *DistrictRepo) Delete(id int64) error {
	err := repo.db.Where("id = ?", id).Delete(&models.District{}).Error
	return err
}
