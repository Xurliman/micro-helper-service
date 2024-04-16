package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type BankBranchRepoInterface interface {
	Create(bankBranch models.BankBranch) (models.BankBranch, error)
	Get(id uuid.UUID) (models.BankBranch, error)
	Update(bankBranch models.BankBranch) (models.BankBranch, error)
	GetByCode(code string) (models.BankBranch, error)
	Delete(id uuid.UUID) error
}

type BankBranchCaseInterface interface {
	Create(bankBranch models.BankBranch) (models.BankBranch, error)
	Get(id uuid.UUID) (models.BankBranch, error)
	Update(bankBranch models.BankBranch) (models.BankBranch, error)
	Delete(id uuid.UUID) error
}
