package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type PassportTypeRepoInterface interface {
	Create(passportType models.PassportType) (models.PassportType, error)
	Get(id uuid.UUID) (models.PassportType, error)
	Update(passportType models.PassportType) (models.PassportType, error)
	GetByCode(code string) (models.PassportType, error)
	Delete(id uuid.UUID) error
}

type PassportTypeCaseInterface interface {
	Create(passportType models.PassportType) (models.PassportType, error)
	Get(id uuid.UUID) (models.PassportType, error)
	Update(passportType models.PassportType) (models.PassportType, error)
	Delete(id uuid.UUID) error
}
