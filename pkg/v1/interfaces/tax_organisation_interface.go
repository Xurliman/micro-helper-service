package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type TaxOrganisationRepoInterface interface {
	Create(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error)
	Get(id uuid.UUID) (models.TaxOrganisation, error)
	Update(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error)
	GetByCode(code string) (models.TaxOrganisation, error)
	Delete(id uuid.UUID) error
}

type TaxOrganisationCaseInterface interface {
	Create(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error)
	Get(id uuid.UUID) (models.TaxOrganisation, error)
	Update(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error)
	Delete(id uuid.UUID) error
}
