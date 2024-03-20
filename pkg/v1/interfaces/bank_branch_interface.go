package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type BankBranchRepoInterface interface {
	Create(bankBranch models.BankBranch) (models.BankBranch, error)
	Get(id int64) (models.BankBranch, error)
	Update(bankBranch models.BankBranch) (models.BankBranch, error)
	GetByCode(code string) (models.BankBranch, error)
	Delete(id int64) error
}

type BankBranchCaseInterface interface {
	Create(bankBranch models.BankBranch) (models.BankBranch, error)
	Get(id int64) (models.BankBranch, error)
	Update(bankBranch models.BankBranch) (models.BankBranch, error)
	Delete(id int64) error
}
