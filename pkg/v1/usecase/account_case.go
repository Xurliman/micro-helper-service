package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountCase struct {
	repo interfaces.AccountRepoInterface
}

func NewAccount(repo interfaces.AccountRepoInterface) interfaces.AccountCaseInterface {
	return &AccountCase{repo: repo}
}

func (accountCase AccountCase) Create(account models.Account) (models.Account, error) {
	if _, err := accountCase.repo.GetByCode(account.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Account{}, errors.New("the code has already been taken")
	}

	return accountCase.repo.Create(account)
}

func (accountCase AccountCase) Get(id uuid.UUID) (models.Account, error) {
	var account models.Account
	var err error

	account, err = accountCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Account{}, errors.New("no such account with the id supplied")
		}

		return models.Account{}, err
	}
	return account, nil
}

func (accountCase AccountCase) Update(updateAccount models.Account) (models.Account, error) {
	var account models.Account
	var err error

	account, err = accountCase.repo.Get(updateAccount.ID)
	if err != nil {
		return models.Account{}, err
	}

	if account.Code != updateAccount.Code {
		return models.Account{}, errors.New("code cannot be changed")
	}

	account, err = accountCase.repo.Update(updateAccount)
	if err != nil {
		return models.Account{}, err
	}
	return account, err
}

func (accountCase AccountCase) Delete(id uuid.UUID) error {
	var err error

	_, err = accountCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = accountCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
