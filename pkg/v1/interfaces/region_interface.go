package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type RegionRepoInterface interface {
	Create(region models.Region) (models.Region, error)
	Get(id uuid.UUID) (models.Region, error)
	Update(region models.Region) (models.Region, error)
	GetByCode(code string) (models.Region, error)
	Delete(id uuid.UUID) error
}

type RegionCaseInterface interface {
	Create(region models.Region) (models.Region, error)
	Get(id uuid.UUID) (models.Region, error)
	Update(region models.Region) (models.Region, error)
	Delete(id uuid.UUID) error
}
