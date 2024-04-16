package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type NationalEconomySectorNewRepoInterface interface {
    Create(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error)
    Get(id uuid.UUID) (models.NationalEconomySectorNew, error)
    Update(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error)
    GetByCode(code string) (models.NationalEconomySectorNew, error)
    Delete(id uuid.UUID) error
}

type NationalEconomySectorNewCaseInterface interface {
    Create(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error)
    Get(id uuid.UUID) (models.NationalEconomySectorNew, error)
    Update(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error)
    Delete(id uuid.UUID) error
}