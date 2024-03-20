package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type DistrictRepoInterface interface {
	Create(district models.District) (models.District, error)
	Get(id int64) (models.District, error)
	Update(district models.District) (models.District, error)
	GetByCode(code string) (models.District, error)
	Delete(id int64) error
}

type DistrictCaseInterface interface {
	Create(district models.District) (models.District, error)
	Get(id int64) (models.District, error)
	Update(district models.District) (models.District, error)
	Delete(id int64) error
}
