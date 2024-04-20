package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BorrowerTypeRepo struct {
	db *gorm.DB
}

func NewBorrowerType(db *gorm.DB) interfaces.BorrowerTypeRepoInterface {
	return &BorrowerTypeRepo{db: db}
}

func (repo *BorrowerTypeRepo) Create(borrowerType models.BorrowerType) (models.BorrowerType, error) {
	err := repo.db.Create(&borrowerType).Error
	return borrowerType, err
}

func (repo *BorrowerTypeRepo) Get(id uuid.UUID) (models.BorrowerType, error) {
	var borrowerType models.BorrowerType
	err := repo.db.Where("id = ?", id).First(&borrowerType).Error

	return borrowerType, err
}

func (repo *BorrowerTypeRepo) Update(borrowerType models.BorrowerType) (models.BorrowerType, error) {
	var dbBorrowerType models.BorrowerType
	if err := repo.db.Where("id = ?", borrowerType.ID).First(&borrowerType).Error; err != nil {
		return dbBorrowerType, err
	}

	dbBorrowerType.Code = borrowerType.Code
	dbBorrowerType.Name = borrowerType.Name
	dbBorrowerType.OldCode = borrowerType.OldCode
	dbBorrowerType.OldName = borrowerType.OldName
	dbBorrowerType.ActivationDate = borrowerType.ActivationDate
	dbBorrowerType.DeactivationDate = borrowerType.DeactivationDate
	dbBorrowerType.FlexFinId = borrowerType.FlexFinId
	dbBorrowerType.NameUz = borrowerType.NameUz

	err := repo.db.Save(dbBorrowerType).Error
	return dbBorrowerType, err
}

func (repo *BorrowerTypeRepo) GetByCode(code string) (models.BorrowerType, error) {
	var borrowerType models.BorrowerType
	err := repo.db.Where("code = ?", code).First(&borrowerType).Error

	return borrowerType, err
}

func (repo *BorrowerTypeRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.BorrowerType{}).Error
	return err
}
