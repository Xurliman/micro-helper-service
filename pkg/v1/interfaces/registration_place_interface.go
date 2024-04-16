package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type RegistrationPlaceRepoInterface interface {
	Create(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error)
	Get(id uuid.UUID) (models.RegistrationPlace, error)
	Update(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error)
	GetByCode(code string) (models.RegistrationPlace, error)
	Delete(id uuid.UUID) error
}

type RegistrationPlaceCaseInterface interface {
	Create(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error)
	Get(id uuid.UUID) (models.RegistrationPlace, error)
	Update(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error)
	Delete(id uuid.UUID) error
}
