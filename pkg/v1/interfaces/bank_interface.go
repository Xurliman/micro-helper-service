package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type BankRepoInterface interface {
	Create(bank models.Bank) (models.Bank, error)
	Get(id uuid.UUID) (models.Bank, error)
	Update(bank models.Bank) (models.Bank, error)
	GetByCode(code string) (models.Bank, error)
	Delete(id uuid.UUID) error
}

type BankCaseInterface interface {
	Create(bank models.Bank) (models.Bank, error)
	Get(id uuid.UUID) (models.Bank, error)
	Update(bank models.Bank) (models.Bank, error)
	Delete(id uuid.UUID) error
}
