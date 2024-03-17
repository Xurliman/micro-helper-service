package usecase

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type NationalEconomySectorCase struct {
	repo interfaces.NationalEconomySectorRepoInterface
}

func (nationalEconomySectorCase NationalEconomySectorCase) Create(nationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error) {
	if _, err := nationalEconomySectorCase.repo.GetByCode(strconv.FormatInt(nationalEconomySector.Code, 10)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.NationalEconomySector{}, errors.New("the code has already been taken")
	}

	return nationalEconomySectorCase.repo.Create(nationalEconomySector)
}

func (nationalEconomySectorCase NationalEconomySectorCase) Get(id string) (models.NationalEconomySector, error) {
	var nationalEconomySector models.NationalEconomySector
	var err error

	nationalEconomySector, err = nationalEconomySectorCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.NationalEconomySector{}, errors.New("no such nationalEconomySector with the id supplied")
		}

		return models.NationalEconomySector{}, err
	}
	return nationalEconomySector, nil
}

func (nationalEconomySectorCase NationalEconomySectorCase) Update(updateNationalEconomySector models.NationalEconomySector) (models.NationalEconomySector, error) {
	var nationalEconomySector models.NationalEconomySector
	var err error

	nationalEconomySector, err = nationalEconomySectorCase.repo.Get(strconv.Itoa(int(updateNationalEconomySector.ID)))
	if err != nil {
		return models.NationalEconomySector{}, err
	}

	if nationalEconomySector.Code != updateNationalEconomySector.Code {
		return models.NationalEconomySector{}, errors.New("code cannot be changed")
	}

	nationalEconomySector, err = nationalEconomySectorCase.repo.Update(updateNationalEconomySector)
	if err != nil {
		return models.NationalEconomySector{}, err
	}
	return nationalEconomySector, err
}

func (nationalEconomySectorCase NationalEconomySectorCase) Delete(id string) error {
	var err error

	_, err = nationalEconomySectorCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = nationalEconomySectorCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
