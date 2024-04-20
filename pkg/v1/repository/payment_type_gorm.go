package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentTypeRepo struct {
	db *gorm.DB
}

func NewPaymentType(db *gorm.DB) interfaces.PaymentTypeRepoInterface {
	return &PaymentTypeRepo{db: db}
}

func (repo *PaymentTypeRepo) Create(paymentType models.PaymentType) (models.PaymentType, error) {
	err := repo.db.Create(&paymentType).Error
	return paymentType, err
}

func (repo *PaymentTypeRepo) Get(id uuid.UUID) (models.PaymentType, error) {
	var paymentType models.PaymentType
	err := repo.db.Where("id = ?", id).First(&paymentType).Error

	return paymentType, err
}

func (repo *PaymentTypeRepo) Update(paymentType models.PaymentType) (models.PaymentType, error) {
	var dbPaymentType models.PaymentType
	if err := repo.db.Where("id = ?", paymentType.ID).First(&paymentType).Error; err != nil {
		return dbPaymentType, err
	}

	dbPaymentType.Code = paymentType.Code
	dbPaymentType.Name = paymentType.Name
	dbPaymentType.ActivationDate = paymentType.ActivationDate
	dbPaymentType.DeactivationDate = paymentType.DeactivationDate
	dbPaymentType.OldCode = paymentType.OldCode
	dbPaymentType.OldName = paymentType.OldName
	dbPaymentType.NameUz = paymentType.NameUz
	dbPaymentType.FlexFinId = paymentType.FlexFinId

	err := repo.db.Save(dbPaymentType).Error
	return dbPaymentType, err
}

func (repo *PaymentTypeRepo) GetByCode(code string) (models.PaymentType, error) {
	var paymentType models.PaymentType
	err := repo.db.Where("code = ?", code).First(&paymentType).Error

	return paymentType, err
}

func (repo *PaymentTypeRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.PaymentType{}).Error
	return err
}
