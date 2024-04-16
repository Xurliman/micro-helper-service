package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type NationalEconomySectorOldRepoInterface interface {
    Create(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error)
    Get(id uuid.UUID) (models.NationalEconomySectorOld, error)
    Update(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error)
    GetByCode(code string) (models.NationalEconomySectorOld, error)
    Delete(id uuid.UUID) error
}

type NationalEconomySectorOldCaseInterface interface {
    Create(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error)
    Get(id uuid.UUID) (models.NationalEconomySectorOld, error)
    Update(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error)
    Delete(id uuid.UUID) error
}