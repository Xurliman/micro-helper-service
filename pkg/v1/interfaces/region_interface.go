package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type RegionRepoInterface interface {
	Create(region models.Region) (models.Region, error)
	Get(id string) (models.Region, error)
	Update(region models.Region) (models.Region, error)
	GetByCode(code string) (models.Region, error)
	Delete(id string) error
}

type RegionCaseInterface interface {
	Create(region models.Region) (models.Region, error)
	Get(id string) (models.Region, error)
	Update(region models.Region) (models.Region, error)
	Delete(id string) error
}
