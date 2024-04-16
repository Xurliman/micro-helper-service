package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type ResidencyTypeRepoInterface interface {
    Create(residencyType models.ResidencyType) (models.ResidencyType, error)
    Get(id uuid.UUID) (models.ResidencyType, error)
    Update(residencyType models.ResidencyType) (models.ResidencyType, error)
    GetByCode(code string) (models.ResidencyType, error)
    Delete(id uuid.UUID) error
}

type ResidencyTypeCaseInterface interface {
    Create(residencyType models.ResidencyType) (models.ResidencyType, error)
    Get(id uuid.UUID) (models.ResidencyType, error)
    Update(residencyType models.ResidencyType) (models.ResidencyType, error)
    Delete(id uuid.UUID) error
}