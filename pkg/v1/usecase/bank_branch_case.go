package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankBranchCase struct {
	repo interfaces.BankBranchRepoInterface
}

func NewBankBranch(repo interfaces.BankBranchRepoInterface) interfaces.BankBranchCaseInterface {
	return &BankBranchCase{repo: repo}
}

func (bankBranchCase BankBranchCase) Create(bankBranch models.BankBranch) (models.BankBranch, error) {
	if _, err := bankBranchCase.repo.GetByCode(bankBranch.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.BankBranch{}, errors.New("the code has already been taken")
	}

	return bankBranchCase.repo.Create(bankBranch)
}

func (bankBranchCase BankBranchCase) Get(id uuid.UUID) (models.BankBranch, error) {
	var bankBranch models.BankBranch
	var err error

	bankBranch, err = bankBranchCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.BankBranch{}, errors.New("no such bankBranch with the id supplied")
		}

		return models.BankBranch{}, err
	}
	return bankBranch, nil
}

func (bankBranchCase BankBranchCase) Update(updateBankBranch models.BankBranch) (models.BankBranch, error) {
	var bankBranch models.BankBranch
	var err error

	bankBranch, err = bankBranchCase.repo.Get(updateBankBranch.ID)
	if err != nil {
		return models.BankBranch{}, err
	}

	if bankBranch.Code != updateBankBranch.Code {
		return models.BankBranch{}, errors.New("code cannot be changed")
	}

	bankBranch, err = bankBranchCase.repo.Update(updateBankBranch)
	if err != nil {
		return models.BankBranch{}, err
	}
	return bankBranch, err
}

func (bankBranchCase BankBranchCase) Delete(id uuid.UUID) error {
	var err error

	_, err = bankBranchCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = bankBranchCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
