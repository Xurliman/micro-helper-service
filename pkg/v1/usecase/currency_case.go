package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CurrencyCase struct {
	repo interfaces.CurrencyRepoInterface
}

func NewCurrency(repo interfaces.CurrencyRepoInterface) interfaces.CurrencyCaseInterface {
	return &CurrencyCase{repo: repo}
}

func (currencyCase *CurrencyCase) Create(currency models.Currency) (models.Currency, error) {
	if _, err := currencyCase.repo.GetByCode(currency.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Currency{}, errors.New("the code has already been taken")
	}

	return currencyCase.repo.Create(currency)
}

func (currencyCase *CurrencyCase) Get(id uuid.UUID) (models.Currency, error) {
	var currency models.Currency
	var err error

	currency, err = currencyCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Currency{}, errors.New("no such currency with the id supplied")
		}

		return models.Currency{}, err
	}
	return currency, nil
}

func (currencyCase *CurrencyCase) Update(updateCurrency models.Currency) (models.Currency, error) {
	var currency models.Currency
	var err error

	currency, err = currencyCase.repo.Get(updateCurrency.ID)
	if err != nil {
		return models.Currency{}, err
	}

	if currency.Code != updateCurrency.Code {
		return models.Currency{}, errors.New("code cannot be changed")
	}

	currency, err = currencyCase.repo.Update(updateCurrency)
	if err != nil {
		return models.Currency{}, err
	}
	return currency, err
}

func (currencyCase *CurrencyCase) Delete(id uuid.UUID) error {
	var err error

	_, err = currencyCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = currencyCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
