package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type CurrencyRepoInterface interface {
	Create(currency models.Currency) (models.Currency, error)
	Get(id uuid.UUID) (models.Currency, error)
	Update(currency models.Currency) (models.Currency, error)
	GetByCode(code string) (models.Currency, error)
	Delete(id uuid.UUID) error
}

type CurrencyCaseInterface interface {
	Create(currency models.Currency) (models.Currency, error)
	Get(id uuid.UUID) (models.Currency, error)
	Update(currency models.Currency) (models.Currency, error)
	Delete(id uuid.UUID) error
}
