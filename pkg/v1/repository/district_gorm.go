package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DistrictRepo struct {
	db *gorm.DB
}

func NewDistrict(db *gorm.DB) interfaces.DistrictRepoInterface {
	return &DistrictRepo{db: db}
}

func (repo *DistrictRepo) Create(district models.District) (models.District, error) {
	err := repo.db.Create(&district).Error
	return district, err
}

func (repo *DistrictRepo) Get(id uuid.UUID) (models.District, error) {
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
	dbDistrict.RegionId = district.RegionId
	dbDistrict.OldCode = district.OldCode
	dbDistrict.OldName = district.OldName
	dbDistrict.ActivationDate = district.ActivationDate
	dbDistrict.DeactivationDate = district.DeactivationDate
	dbDistrict.FlexFinId = district.FlexFinId

	err := repo.db.Save(dbDistrict).Error
	return dbDistrict, err
}

func (repo *DistrictRepo) GetByCode(code string) (models.District, error) {
	var district models.District
	err := repo.db.Where("code = ?", code).First(&district).Error

	return district, err
}

func (repo *DistrictRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.District{}).Error
	return err
}
