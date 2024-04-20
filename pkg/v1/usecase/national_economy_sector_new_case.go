package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NationalEconomySectorNewCase struct {
	repo interfaces.NationalEconomySectorNewRepoInterface
}

func NewNationalEconomySectorNew(repo interfaces.NationalEconomySectorNewRepoInterface) interfaces.NationalEconomySectorNewCaseInterface {
	return &NationalEconomySectorNewCase{repo: repo}
}

func (nationalEconomySectorNewCase NationalEconomySectorNewCase) Create(nationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error) {
	if _, err := nationalEconomySectorNewCase.repo.GetByCode(nationalEconomySectorNew.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.NationalEconomySectorNew{}, errors.New("the code has already been taken")
	}

	return nationalEconomySectorNewCase.repo.Create(nationalEconomySectorNew)
}

func (nationalEconomySectorNewCase NationalEconomySectorNewCase) Get(id uuid.UUID) (models.NationalEconomySectorNew, error) {
	var nationalEconomySectorNew models.NationalEconomySectorNew
	var err error

	nationalEconomySectorNew, err = nationalEconomySectorNewCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.NationalEconomySectorNew{}, errors.New("no such nationalEconomySectorNew with the id supplied")
		}

		return models.NationalEconomySectorNew{}, err
	}
	return nationalEconomySectorNew, nil
}

func (nationalEconomySectorNewCase NationalEconomySectorNewCase) Update(updateNationalEconomySectorNew models.NationalEconomySectorNew) (models.NationalEconomySectorNew, error) {
	var nationalEconomySectorNew models.NationalEconomySectorNew
	var err error

	nationalEconomySectorNew, err = nationalEconomySectorNewCase.repo.Get(updateNationalEconomySectorNew.ID)
	if err != nil {
		return models.NationalEconomySectorNew{}, err
	}

	if nationalEconomySectorNew.Code != updateNationalEconomySectorNew.Code {
		return models.NationalEconomySectorNew{}, errors.New("code cannot be changed")
	}

	nationalEconomySectorNew, err = nationalEconomySectorNewCase.repo.Update(updateNationalEconomySectorNew)
	if err != nil {
		return models.NationalEconomySectorNew{}, err
	}
	return nationalEconomySectorNew, err
}

func (nationalEconomySectorNewCase NationalEconomySectorNewCase) Delete(id uuid.UUID) error {
	var err error

	_, err = nationalEconomySectorNewCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = nationalEconomySectorNewCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
