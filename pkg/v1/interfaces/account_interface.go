package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type AccountRepoInterface interface {
	Create(account models.Account) (models.Account, error)
	Get(id uuid.UUID) (models.Account, error)
	Update(account models.Account) (models.Account, error)
	GetByCode(code string) (models.Account, error)
	Delete(id uuid.UUID) error
}

type AccountCaseInterface interface {
	Create(account models.Account) (models.Account, error)
	Get(id uuid.UUID) (models.Account, error)
	Update(account models.Account) (models.Account, error)
	Delete(id uuid.UUID) error
}
