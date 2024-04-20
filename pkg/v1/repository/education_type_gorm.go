package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EducationTypeRepo struct {
	db *gorm.DB
}

func NewEducationType(db *gorm.DB) interfaces.EducationTypeRepoInterface {
	return &EducationTypeRepo{db: db}
}

func (repo *EducationTypeRepo) Create(educationType models.EducationType) (models.EducationType, error) {
	err := repo.db.Create(&educationType).Error
	return educationType, err
}

func (repo *EducationTypeRepo) Get(id uuid.UUID) (models.EducationType, error) {
	var educationType models.EducationType
	err := repo.db.Where("id = ?", id).First(&educationType).Error

	return educationType, err
}

func (repo *EducationTypeRepo) Update(educationType models.EducationType) (models.EducationType, error) {
	var dbEducationType models.EducationType
	if err := repo.db.Where("id = ?", educationType.ID).First(&educationType).Error; err != nil {
		return dbEducationType, err
	}

	dbEducationType.Code = educationType.Code
	dbEducationType.Name = educationType.Name
	dbEducationType.FlexFinId = educationType.FlexFinId

	err := repo.db.Save(dbEducationType).Error
	return dbEducationType, err
}

func (repo *EducationTypeRepo) GetByCode(code string) (models.EducationType, error) {
	var educationType models.EducationType
	err := repo.db.Where("code = ?", code).First(&educationType).Error

	return educationType, err
}

func (repo *EducationTypeRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.EducationType{}).Error
	return err
}
