package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaxOrganisationCase struct {
	repo interfaces.TaxOrganisationRepoInterface
}

func NewTaxOrganisation(repo interfaces.TaxOrganisationRepoInterface) interfaces.TaxOrganisationCaseInterface {
	return &TaxOrganisationCase{repo: repo}
}

func (taxOrganisationCase TaxOrganisationCase) Create(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error) {
	if _, err := taxOrganisationCase.repo.GetByCode(taxOrganisation.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.TaxOrganisation{}, errors.New("the code has already been taken")
	}

	return taxOrganisationCase.repo.Create(taxOrganisation)
}

func (taxOrganisationCase TaxOrganisationCase) Get(id uuid.UUID) (models.TaxOrganisation, error) {
	var taxOrganisation models.TaxOrganisation
	var err error

	taxOrganisation, err = taxOrganisationCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.TaxOrganisation{}, errors.New("no such taxOrganisation with the id supplied")
		}

		return models.TaxOrganisation{}, err
	}
	return taxOrganisation, nil
}

func (taxOrganisationCase TaxOrganisationCase) Update(updateTaxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error) {
	var taxOrganisation models.TaxOrganisation
	var err error

	taxOrganisation, err = taxOrganisationCase.repo.Get(updateTaxOrganisation.ID)
	if err != nil {
		return models.TaxOrganisation{}, err
	}

	if taxOrganisation.Code != updateTaxOrganisation.Code {
		return models.TaxOrganisation{}, errors.New("code cannot be changed")
	}

	taxOrganisation, err = taxOrganisationCase.repo.Update(updateTaxOrganisation)
	if err != nil {
		return models.TaxOrganisation{}, err
	}
	return taxOrganisation, err
}

func (taxOrganisationCase TaxOrganisationCase) Delete(id uuid.UUID) error {
	var err error

	_, err = taxOrganisationCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = taxOrganisationCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
