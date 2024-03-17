package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type BankRepoInterface interface {
	Create(bank models.Bank) (models.Bank, error)
	Get(id string) (models.Bank, error)
	Update(bank models.Bank) (models.Bank, error)
	GetByCode(code string) (models.Bank, error)
	Delete(id string) error
}

type BankCaseInterface interface {
	Create(bank models.Bank) (models.Bank, error)
	Get(id string) (models.Bank, error)
	Update(bank models.Bank) (models.Bank, error)
	Delete(id string) error
}
