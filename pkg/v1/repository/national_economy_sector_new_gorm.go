package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NationalEconomySectorNewRepo struct {
	db *gorm.DB
}

func (repo *NationalEconomySectorNewRepo) Create(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error) {
	err := repo.db.Create(&nationalEconomySectorNew).Error
	return nationalEconomySectorNew, err
}

func (repo *NationalEconomySectorNewRepo) Get(id uuid.UUID) (models.NationalEconomySectorNew, error) {
	var nationalEconomySectorNew models.NationalEconomySectorNew
	err := repo.db.Where("id = ?", id).First(&nationalEconomySectorNew).Error

	return nationalEconomySectorNew, err
}

func (repo *NationalEconomySectorNewRepo) Update(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error) {
	var dbNationalEconomySectorNew models.NationalEconomySectorNew
	if err := repo.db.Where("id = ?", nationalEconomySectorNew.ID).First(&nationalEconomySectorNew).Error; err != nil {
		return dbNationalEconomySectorNew, err
	}

	dbNationalEconomySectorNew.Code = nationalEconomySectorNew.Code
	dbNationalEconomySectorNew.Name = nationalEconomySectorNew.Name
	dbNationalEconomySectorNew.NameUz = nationalEconomySectorNew.NameUz
	dbNationalEconomySectorNew.Section = nationalEconomySectorNew.Section
	dbNationalEconomySectorNew.SectionCode = nationalEconomySectorNew.SectionCode
	dbNationalEconomySectorNew.SectionName = nationalEconomySectorNew.SectionName
	dbNationalEconomySectorNew.SectionNameUz = nationalEconomySectorNew.SectionNameUz
	dbNationalEconomySectorNew.GroupCode = nationalEconomySectorNew.GroupCode
	dbNationalEconomySectorNew.GroupName = nationalEconomySectorNew.GroupName
	dbNationalEconomySectorNew.GroupNameUz = nationalEconomySectorNew.GroupNameUz
	dbNationalEconomySectorNew.Group1Code = nationalEconomySectorNew.Group1Code
	dbNationalEconomySectorNew.Group1Name = nationalEconomySectorNew.Group1Name
	dbNationalEconomySectorNew.Group1NameUz = nationalEconomySectorNew.Group1NameUz
	dbNationalEconomySectorNew.Group2Code = nationalEconomySectorNew.Group2Code
	dbNationalEconomySectorNew.Group2Name = nationalEconomySectorNew.Group2Name
	dbNationalEconomySectorNew.Group2NameUz = nationalEconomySectorNew.Group2NameUz
	dbNationalEconomySectorNew.ActivationDate = nationalEconomySectorNew.ActivationDate
	dbNationalEconomySectorNew.DeactivationDate = nationalEconomySectorNew.DeactivationDate
	dbNationalEconomySectorNew.CBUReferenceKey = nationalEconomySectorNew.CBUReferenceKey
	dbNationalEconomySectorNew.FlexFinId = nationalEconomySectorNew.FlexFinId

	err := repo.db.Save(dbNationalEconomySectorNew).Error
	return dbNationalEconomySectorNew, err
}

func (repo *NationalEconomySectorNewRepo) GetByCode(code string) (models.NationalEconomySectorNew, error) {
	var nationalEconomySectorNew models.NationalEconomySectorNew
	err := repo.db.Where("code = ?", code).First(&nationalEconomySectorNew).Error

	return nationalEconomySectorNew, err
}

func (repo *NationalEconomySectorNewRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.NationalEconomySectorNew{}).Error
	return err
}
