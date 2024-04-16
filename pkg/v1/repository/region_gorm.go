package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegionRepo struct {
	db *gorm.DB
}

func (repo *RegionRepo) Create(region models.Region) (models.Region, error) {
	err := repo.db.Create(&region).Error
	return region, err
}

func (repo *RegionRepo) Get(id uuid.UUID) (models.Region, error) {
	var region models.Region
	err := repo.db.Where("id = ?", id).First(&region).Error

	return region, err
}

func (repo *RegionRepo) Update(region models.Region) (models.Region, error) {
	var dbRegion models.Region
	if err := repo.db.Where("id = ?", region.ID).First(&region).Error; err != nil {
		return dbRegion, err
	}

	dbRegion.Code = region.Code
	dbRegion.Name = region.Name
	dbRegion.CountryId = region.CountryId
	dbRegion.ActivationDate = region.ActivationDate
	dbRegion.DeactivationDate = region.DeactivationDate
	dbRegion.OldCode = region.OldCode
	dbRegion.OldName = region.OldName
	dbRegion.FlexFinId = region.FlexFinId

	err := repo.db.Save(dbRegion).Error
	return dbRegion, err
}

func (repo *RegionRepo) GetByCode(code string) (models.Region, error) {
	var region models.Region
	err := repo.db.Where("code = ?", code).First(&region).Error

	return region, err
}

func (repo *RegionRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.Region{}).Error
	return err
}
