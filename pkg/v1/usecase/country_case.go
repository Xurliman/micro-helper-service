package usecase

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type CountryCase struct {
	repo interfaces.CountryRepoInterface
}

func (countryCase *CountryCase) Create(country models.Country) (models.Country, error) {
	if _, err := countryCase.repo.GetByCode(strconv.FormatInt(country.Code, 10)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Country{}, errors.New("the code has already been taken")
	}

	return countryCase.repo.Create(country)
}

func (countryCase *CountryCase) Get(id int64) (models.Country, error) {
	var country models.Country
	var err error

	country, err = countryCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Country{}, errors.New("no such country with the id supplied")
		}

		return models.Country{}, err
	}
	return country, nil
}

func (countryCase *CountryCase) Update(updateCountry models.Country) (models.Country, error) {
	var country models.Country
	var err error

	country, err = countryCase.repo.Get(int64(updateCountry.ID))
	if err != nil {
		return models.Country{}, err
	}

	if country.Code != updateCountry.Code {
		return models.Country{}, errors.New("code cannot be changed")
	}

	country, err = countryCase.repo.Update(updateCountry)
	if err != nil {
		return models.Country{}, err
	}
	return country, err
}

func (countryCase *CountryCase) Delete(id int64) error {
	var err error

	_, err = countryCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = countryCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
