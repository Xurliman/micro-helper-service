package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type CountryRepoInterface interface {
	Create(country models.Country) (models.Country, error)
	Get(id int64) (models.Country, error)
	Update(country models.Country) (models.Country, error)
	GetByCode(code string) (models.Country, error)
	Delete(id int64) error
}

type CountryCaseInterface interface {
	Create(country models.Country) (models.Country, error)
	Get(id int64) (models.Country, error)
	Update(country models.Country) (models.Country, error)
	Delete(id int64) error
}
