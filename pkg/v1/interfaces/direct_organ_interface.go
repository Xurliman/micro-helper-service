package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type DirectOrganRepoInterface interface {
	Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Get(id uuid.UUID) (models.DirectOrgan, error)
	Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	GetByCode(code string) (models.DirectOrgan, error)
	Delete(id uuid.UUID) error
}

type DirectOrganCaseInterface interface {
	Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Get(id uuid.UUID) (models.DirectOrgan, error)
	Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Delete(id uuid.UUID) error
}
