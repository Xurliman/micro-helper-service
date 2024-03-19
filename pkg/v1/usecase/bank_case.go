package usecase

import (
	"errors"
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"gorm.io/gorm"
	"strconv"
)

type BankCase struct {
	repo interfaces.BankRepoInterface
}

func NewBank(repo interfaces.BankRepoInterface) interfaces.BankCaseInterface {
	return &BankCase{repo: repo}
}

func (bankCase *BankCase) Create(bank models.Bank) (models.Bank, error) {
	if _, err := bankCase.repo.GetByCode(strconv.FormatInt(bank.Code, 10)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Bank{}, errors.New("the code has already been taken")
	}

	return bankCase.repo.Create(bank)
}

func (bankCase *BankCase) Get(id string) (models.Bank, error) {
	var bank models.Bank
	var err error

	bank, err = bankCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Bank{}, errors.New("no such bank with the id supplied")
		}

		return models.Bank{}, err
	}
	return bank, nil
}

func (bankCase *BankCase) Update(updateBank models.Bank) (models.Bank, error) {
	var bank models.Bank
	var err error

	bank, err = bankCase.repo.Get(strconv.Itoa(int(updateBank.ID)))
	if err != nil {
		return models.Bank{}, err
	}

	if bank.Code != updateBank.Code {
		return models.Bank{}, errors.New("code cannot be changed")
	}

	bank, err = bankCase.repo.Update(updateBank)
	if err != nil {
		return models.Bank{}, err
	}
	return bank, err
}

func (bankCase *BankCase) Delete(id string) error {
	var err error

	_, err = bankCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = bankCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
