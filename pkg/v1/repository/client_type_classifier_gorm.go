package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
)

type ClientTypeClassifierRepo struct {
	db *gorm.DB
}

func (repo *ClientTypeClassifierRepo) Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error) {
	err := repo.db.Create(&classifier).Error
	return classifier, err
}

func (repo *ClientTypeClassifierRepo) Get(id string) (models.ClientTypeClassifier, error) {
	var classifier models.ClientTypeClassifier
	err := repo.db.Where("id = ?", id).First(&classifier).Error

	return classifier, err
}

func (repo *ClientTypeClassifierRepo) Update(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error) {
	var dbClassifier models.ClientTypeClassifier
	if err := repo.db.Where("id = ?", classifier.ID).First(&classifier).Error; err != nil {
		return dbClassifier, err
	}

	dbClassifier.Code = classifier.Code
	dbClassifier.Name = classifier.Name
	dbClassifier.ShortName = classifier.ShortName
	dbClassifier.ClientType = classifier.ClientType
	dbClassifier.ActivationDate = classifier.ActivationDate
	dbClassifier.DeactivationDate = classifier.DeactivationDate
	dbClassifier.CBUReferenceKey = classifier.CBUReferenceKey
	dbClassifier.OldCode = classifier.OldCode
	dbClassifier.OldName = classifier.OldName

	err := repo.db.Save(dbClassifier).Error
	return dbClassifier, err
}

func (repo *ClientTypeClassifierRepo) GetByCode(code string) (models.ClientTypeClassifier, error) {
	var classifier models.ClientTypeClassifier
	err := repo.db.Where("code = ?", code).First(&classifier).Error

	return classifier, err
}

func (repo *ClientTypeClassifierRepo) Delete(id string) error {
	err := repo.db.Where("id = ?", id).Delete(&models.ClientTypeClassifier{}).Error
	return err
}
