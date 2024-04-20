package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BorrowerTypeCase struct {
	repo interfaces.BorrowerTypeRepoInterface
}

func NewBorrowerType(repo interfaces.BorrowerTypeRepoInterface) interfaces.BorrowerTypeCaseInterface {
	return &BorrowerTypeCase{repo: repo}
}

func (borrowerTypeCase BorrowerTypeCase) Create(borrowerType models.BorrowerType) (models.BorrowerType, error) {
	if _, err := borrowerTypeCase.repo.GetByCode(borrowerType.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.BorrowerType{}, errors.New("the code has already been taken")
	}

	return borrowerTypeCase.repo.Create(borrowerType)
}

func (borrowerTypeCase BorrowerTypeCase) Get(id uuid.UUID) (models.BorrowerType, error) {
	var borrowerType models.BorrowerType
	var err error

	borrowerType, err = borrowerTypeCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.BorrowerType{}, errors.New("no such borrowerType with the id supplied")
		}

		return models.BorrowerType{}, err
	}
	return borrowerType, nil
}

func (borrowerTypeCase BorrowerTypeCase) Update(updateBorrowerType models.BorrowerType) (models.BorrowerType, error) {
	var borrowerType models.BorrowerType
	var err error

	borrowerType, err = borrowerTypeCase.repo.Get(updateBorrowerType.ID)
	if err != nil {
		return models.BorrowerType{}, err
	}

	if borrowerType.Code != updateBorrowerType.Code {
		return models.BorrowerType{}, errors.New("code cannot be changed")
	}

	borrowerType, err = borrowerTypeCase.repo.Update(updateBorrowerType)
	if err != nil {
		return models.BorrowerType{}, err
	}
	return borrowerType, err
}

func (borrowerTypeCase BorrowerTypeCase) Delete(id uuid.UUID) error {
	var err error

	_, err = borrowerTypeCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = borrowerTypeCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
