package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type NationalEconomySectorRepoInterface interface {
	Create(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error)
	Get(id string) (models.NationalEconomySector, error)
	Update(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error)
	GetByCode(code string) (models.NationalEconomySector, error)
	Delete(id string) error
}

type NationalEconomySectorCaseInterface interface {
	Create(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error)
	Get(id string) (models.NationalEconomySector, error)
	Update(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error)
	Delete(id string) error
}
