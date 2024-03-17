package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type PaymentTypeRepoInterface interface {
	Create(paymentType models.PaymentType) (models.PaymentType, error)
	Get(id string) (models.PaymentType, error)
	Update(paymentType models.PaymentType) (models.PaymentType, error)
	GetByCode(code string) (models.PaymentType, error)
	Delete(id string) error
}

type PaymentTypeCaseInterface interface {
	Create(paymentType models.PaymentType) (models.PaymentType, error)
	Get(id string) (models.PaymentType, error)
	Update(paymentType models.PaymentType) (models.PaymentType, error)
	Delete(id string) error
}
