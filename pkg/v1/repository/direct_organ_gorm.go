package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DirectOrganRepo struct {
	db *gorm.DB
}

func NewDirectOrgan(db *gorm.DB) interfaces.DirectOrganRepoInterface {
	return &DirectOrganRepo{db: db}
}

func (repo *DirectOrganRepo) Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error) {
	err := repo.db.Create(&directOrgan).Error
	return directOrgan, err
}

func (repo *DirectOrganRepo) Get(id uuid.UUID) (models.DirectOrgan, error) {
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
	dbDirectOrgan.CBUCode = directOrgan.CBUCode
	dbDirectOrgan.ActivationDate = directOrgan.ActivationDate
	dbDirectOrgan.DeactivationDate = directOrgan.DeactivationDate
	dbDirectOrgan.FlexFinId = directOrgan.FlexFinId

	err := repo.db.Save(dbDirectOrgan).Error
	return dbDirectOrgan, err
}

func (repo *DirectOrganRepo) GetByCode(code string) (models.DirectOrgan, error) {
	var directOrgan models.DirectOrgan
	err := repo.db.Where("code = ?", code).First(&directOrgan).Error

	return directOrgan, err
}

func (repo *DirectOrganRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.DirectOrgan{}).Error
	return err
}
