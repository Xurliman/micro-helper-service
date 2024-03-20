package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type RegionRepoInterface interface {
	Create(region models.Region) (models.Region, error)
	Get(id int64) (models.Region, error)
	Update(region models.Region) (models.Region, error)
	GetByCode(code string) (models.Region, error)
	Delete(id int64) error
}

type RegionCaseInterface interface {
	Create(region models.Region) (models.Region, error)
	Get(id int64) (models.Region, error)
	Update(region models.Region) (models.Region, error)
	Delete(id int64) error
}
