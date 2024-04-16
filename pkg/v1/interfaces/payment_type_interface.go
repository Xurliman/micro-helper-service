package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type PaymentTypeRepoInterface interface {
	Create(paymentType models.PaymentType) (models.PaymentType, error)
	Get(id uuid.UUID) (models.PaymentType, error)
	Update(paymentType models.PaymentType) (models.PaymentType, error)
	GetByCode(code string) (models.PaymentType, error)
	Delete(id uuid.UUID) error
}

type PaymentTypeCaseInterface interface {
	Create(paymentType models.PaymentType) (models.PaymentType, error)
	Get(id uuid.UUID) (models.PaymentType, error)
	Update(paymentType models.PaymentType) (models.PaymentType, error)
	Delete(id uuid.UUID) error
}
