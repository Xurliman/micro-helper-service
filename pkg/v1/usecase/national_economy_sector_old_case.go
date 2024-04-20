package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NationalEconomySectorOldCase struct {
	repo interfaces.NationalEconomySectorOldRepoInterface
}

func NewNationalEconomySectorOld(repo interfaces.NationalEconomySectorOldRepoInterface) interfaces.NationalEconomySectorOldCaseInterface {
	return &NationalEconomySectorOldCase{repo: repo}
}

func (nationalEconomySectorOldCase NationalEconomySectorOldCase) Create(nationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error) {
	if _, err := nationalEconomySectorOldCase.repo.GetByCode(nationalEconomySectorOld.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.NationalEconomySectorOld{}, errors.New("the code has already been taken")
	}

	return nationalEconomySectorOldCase.repo.Create(nationalEconomySectorOld)
}

func (nationalEconomySectorOldCase NationalEconomySectorOldCase) Get(id uuid.UUID) (models.NationalEconomySectorOld, error) {
	var nationalEconomySectorOld models.NationalEconomySectorOld
	var err error

	nationalEconomySectorOld, err = nationalEconomySectorOldCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.NationalEconomySectorOld{}, errors.New("no such nationalEconomySectorOld with the id supplied")
		}

		return models.NationalEconomySectorOld{}, err
	}
	return nationalEconomySectorOld, nil
}

func (nationalEconomySectorOldCase NationalEconomySectorOldCase) Update(updateNationalEconomySectorOld models.NationalEconomySectorOld) (models.NationalEconomySectorOld, error) {
	var nationalEconomySectorOld models.NationalEconomySectorOld
	var err error

	nationalEconomySectorOld, err = nationalEconomySectorOldCase.repo.Get(updateNationalEconomySectorOld.ID)
	if err != nil {
		return models.NationalEconomySectorOld{}, err
	}

	if nationalEconomySectorOld.Code != updateNationalEconomySectorOld.Code {
		return models.NationalEconomySectorOld{}, errors.New("code cannot be changed")
	}

	nationalEconomySectorOld, err = nationalEconomySectorOldCase.repo.Update(updateNationalEconomySectorOld)
	if err != nil {
		return models.NationalEconomySectorOld{}, err
	}
	return nationalEconomySectorOld, err
}

func (nationalEconomySectorOldCase NationalEconomySectorOldCase) Delete(id uuid.UUID) error {
	var err error

	_, err = nationalEconomySectorOldCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = nationalEconomySectorOldCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
