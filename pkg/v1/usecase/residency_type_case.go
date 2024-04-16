package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResidencyTypeCase struct {
	repo interfaces.ResidencyTypeRepoInterface
}

func (residencyTypeCase ResidencyTypeCase) Create(residencyType models.ResidencyType) (models.ResidencyType, error) {
	if _, err := residencyTypeCase.repo.GetByCode(residencyType.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ResidencyType{}, errors.New("the code has already been taken")
	}

	return residencyTypeCase.repo.Create(residencyType)
}

func (residencyTypeCase ResidencyTypeCase) Get(id uuid.UUID) (models.ResidencyType, error) {
	var residencyType models.ResidencyType
	var err error

	residencyType, err = residencyTypeCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ResidencyType{}, errors.New("no such residencyType with the id supplied")
		}

		return models.ResidencyType{}, err
	}
	return residencyType, nil
}

func (residencyTypeCase ResidencyTypeCase) Update(updateResidencyType models.ResidencyType) (models.ResidencyType, error) {
	var residencyType models.ResidencyType
	var err error

	residencyType, err = residencyTypeCase.repo.Get(updateResidencyType.ID)
	if err != nil {
		return models.ResidencyType{}, err
	}

	if residencyType.Code != updateResidencyType.Code {
		return models.ResidencyType{}, errors.New("code cannot be changed")
	}

	residencyType, err = residencyTypeCase.repo.Update(updateResidencyType)
	if err != nil {
		return models.ResidencyType{}, err
	}
	return residencyType, err
}

func (residencyTypeCase ResidencyTypeCase) Delete(id uuid.UUID) error {
	var err error

	_, err = residencyTypeCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = residencyTypeCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
