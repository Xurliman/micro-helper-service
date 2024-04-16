package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type BorrowerTypeRepoInterface interface {
	Create(borrowerType models.BorrowerType) (models.BorrowerType, error)
	Get(id uuid.UUID) (models.BorrowerType, error)
	Update(borrowerType models.BorrowerType) (models.BorrowerType, error)
	GetByCode(code string) (models.BorrowerType, error)
	Delete(id uuid.UUID) error
}

type BorrowerTypeCaseInterface interface {
	Create(borrowerType models.BorrowerType) (models.BorrowerType, error)
	Get(id uuid.UUID) (models.BorrowerType, error)
	Update(borrowerType models.BorrowerType) (models.BorrowerType, error)
	Delete(id uuid.UUID) error
}
