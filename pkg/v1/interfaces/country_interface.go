package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type CountryRepoInterface interface {
	Create(country models.Country) (models.Country, error)
	Get(id uuid.UUID) (models.Country, error)
	Update(country models.Country) (models.Country, error)
	GetByCode(code string) (models.Country, error)
	Delete(id uuid.UUID) error
}

type CountryCaseInterface interface {
	Create(country models.Country) (models.Country, error)
	Get(id uuid.UUID) (models.Country, error)
	Update(country models.Country) (models.Country, error)
	Delete(id uuid.UUID) error
}
