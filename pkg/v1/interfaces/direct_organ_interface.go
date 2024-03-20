package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type DirectOrganRepoInterface interface {
	Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Get(id int64) (models.DirectOrgan, error)
	Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	GetByCode(code string) (models.DirectOrgan, error)
	Delete(id int64) error
}

type DirectOrganCaseInterface interface {
	Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Get(id int64) (models.DirectOrgan, error)
	Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Delete(id int64) error
}
