package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
)

type DirectOrganRepo struct {
	db *gorm.DB
}

func (repo *DirectOrganRepo) Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error) {
	err := repo.db.Create(&directOrgan).Error
	return directOrgan, err
}

func (repo *DirectOrganRepo) Get(id int64) (models.DirectOrgan, error) {
	var directOrgan models.DirectOrgan
	err := repo.db.Where("id = ?", id).First(&directOrgan).Error

	return directOrgan, err
}

func (repo *DirectOrganRepo) Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error) {
	var dbDirectOrgan models.DirectOrgan
	if err := repo.db.Where("id = ?", directOrgan.ID).First(&directOrgan).Error; err != nil {
		return dbDirectOrgan, err
	}

	dbDirectOrgan.Code = directOrgan.Code
	dbDirectOrgan.Name = directOrgan.Name
	dbDirectOrgan.ShortName = directOrgan.ShortName
	dbDirectOrgan.CrudDates = directOrgan.CrudDates
	dbDirectOrgan.CBUCode = directOrgan.CBUCode

	err := repo.db.Save(dbDirectOrgan).Error
	return dbDirectOrgan, err
}

func (repo *DirectOrganRepo) GetByCode(code string) (models.DirectOrgan, error) {
	var directOrgan models.DirectOrgan
	err := repo.db.Where("code = ?", code).First(&directOrgan).Error

	return directOrgan, err
}

func (repo *DirectOrganRepo) Delete(id int64) error {
	err := repo.db.Where("id = ?", id).Delete(&models.DirectOrgan{}).Error
	return err
}
