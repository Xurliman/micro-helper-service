package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PassportTypeCase struct {
	repo interfaces.PassportTypeRepoInterface
}

func (passportTypeCase PassportTypeCase) Create(passportType models.PassportType) (models.PassportType, error) {
	if _, err := passportTypeCase.repo.GetByCode(passportType.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.PassportType{}, errors.New("the code has already been taken")
	}

	return passportTypeCase.repo.Create(passportType)
}

func (passportTypeCase PassportTypeCase) Get(id uuid.UUID) (models.PassportType, error) {
	var passportType models.PassportType
	var err error

	passportType, err = passportTypeCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.PassportType{}, errors.New("no such passportType with the id supplied")
		}

		return models.PassportType{}, err
	}
	return passportType, nil
}

func (passportTypeCase PassportTypeCase) Update(updatePassportType models.PassportType) (models.PassportType, error) {
	var passportType models.PassportType
	var err error

	passportType, err = passportTypeCase.repo.Get(updatePassportType.ID)
	if err != nil {
		return models.PassportType{}, err
	}

	if passportType.Code != updatePassportType.Code {
		return models.PassportType{}, errors.New("code cannot be changed")
	}

	passportType, err = passportTypeCase.repo.Update(updatePassportType)
	if err != nil {
		return models.PassportType{}, err
	}
	return passportType, err
}

func (passportTypeCase PassportTypeCase) Delete(id uuid.UUID) error {
	var err error

	_, err = passportTypeCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = passportTypeCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
