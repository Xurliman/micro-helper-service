package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegistrationPlaceCase struct {
	repo interfaces.RegistrationPlaceRepoInterface
}

func (registrationPlaceCase RegistrationPlaceCase) Create(registrationPlace models.RegistrationPlace) (models.RegistrationPlace, error) {
	if _, err := registrationPlaceCase.repo.GetByCode(registrationPlace.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.RegistrationPlace{}, errors.New("the code has already been taken")
	}

	return registrationPlaceCase.repo.Create(registrationPlace)
}

func (registrationPlaceCase RegistrationPlaceCase) Get(id uuid.UUID) (models.RegistrationPlace, error) {
	var registrationPlace models.RegistrationPlace
	var err error

	registrationPlace, err = registrationPlaceCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.RegistrationPlace{}, errors.New("no such registrationPlace with the id supplied")
		}

		return models.RegistrationPlace{}, err
	}
	return registrationPlace, nil
}

func (registrationPlaceCase RegistrationPlaceCase) Update(updateRegistrationPlace models.RegistrationPlace) (models.RegistrationPlace, error) {
	var registrationPlace models.RegistrationPlace
	var err error

	registrationPlace, err = registrationPlaceCase.repo.Get(updateRegistrationPlace.ID)
	if err != nil {
		return models.RegistrationPlace{}, err
	}

	if registrationPlace.Code != updateRegistrationPlace.Code {
		return models.RegistrationPlace{}, errors.New("code cannot be changed")
	}

	registrationPlace, err = registrationPlaceCase.repo.Update(updateRegistrationPlace)
	if err != nil {
		return models.RegistrationPlace{}, err
	}
	return registrationPlace, err
}

func (registrationPlaceCase RegistrationPlaceCase) Delete(id uuid.UUID) error {
	var err error

	_, err = registrationPlaceCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = registrationPlaceCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
