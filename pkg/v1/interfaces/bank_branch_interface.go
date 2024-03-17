package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type BankRepoInterBankBranch interface {
	Create(bankBranch models.BankBranch) (models.BankBranch, error)
	Get(id string) (models.BankBranch, error)
	Update(id string, bankBranch models.BankBranch) (models.BankBranch, error)
	GetByCode(code string) (models.BankBranch, error)
	Delete(id string) error
}

type BankCaseInterBankBranch interface {
	Create(bankBranch models.BankBranch) (models.BankBranch, error)
	Get(id string) (models.BankBranch, error)
	Update(id string, bankBranch models.BankBranch) (models.BankBranch, error)
	Delete(id string) error
}
