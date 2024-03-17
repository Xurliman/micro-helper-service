package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type DistrictRepoInterface interface {
	Create(district models.District) (models.District, error)
	Get(id string) (models.District, error)
	Update(district models.District) (models.District, error)
	GetByCode(code string) (models.District, error)
	Delete(id string) error
}

type DistrictCaseInterface interface {
	Create(district models.District) (models.District, error)
	Get(id string) (models.District, error)
	Update(district models.District) (models.District, error)
	Delete(id string) error
}
