package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type DistrictRepoInterface interface {
	Create(district models.District) (models.District, error)
	Get(id uuid.UUID) (models.District, error)
	Update(district models.District) (models.District, error)
	GetByCode(code string) (models.District, error)
	Delete(id uuid.UUID) error
}

type DistrictCaseInterface interface {
	Create(district models.District) (models.District, error)
	Get(id uuid.UUID) (models.District, error)
	Update(district models.District) (models.District, error)
	Delete(id uuid.UUID) error
}
