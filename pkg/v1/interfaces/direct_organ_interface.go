package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type DirectOrganRepoInterface interface {
	Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Get(id string) (models.DirectOrgan, error)
	Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	GetByCode(code string) (models.DirectOrgan, error)
	Delete(id string) error
}

type DirectOrganCaseInterface interface {
	Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Get(id string) (models.DirectOrgan, error)
	Update(directOrgan models.DirectOrgan) (models.DirectOrgan, error)
	Delete(id string) error
}
