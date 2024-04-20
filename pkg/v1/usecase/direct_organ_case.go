package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DirectOrganCase struct {
	repo interfaces.DirectOrganRepoInterface
}

func NewDirectOrgan(repo interfaces.DirectOrganRepoInterface) interfaces.DirectOrganCaseInterface {
	return &DirectOrganCase{repo: repo}
}

func (directOrganCase *DirectOrganCase) Create(directOrgan models.DirectOrgan) (models.DirectOrgan, error) {
	if _, err := directOrganCase.repo.GetByCode(directOrgan.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.DirectOrgan{}, errors.New("the code has already been taken")
	}

	return directOrganCase.repo.Create(directOrgan)
}

func (directOrganCase *DirectOrganCase) Get(id uuid.UUID) (models.DirectOrgan, error) {
	var directOrgan models.DirectOrgan
	var err error

	directOrgan, err = directOrganCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.DirectOrgan{}, errors.New("no such directOrgan with the id supplied")
		}

		return models.DirectOrgan{}, err
	}
	return directOrgan, nil
}

func (directOrganCase *DirectOrganCase) Update(updateDirectOrgan models.DirectOrgan) (models.DirectOrgan, error) {
	var directOrgan models.DirectOrgan
	var err error

	directOrgan, err = directOrganCase.repo.Get(updateDirectOrgan.ID)
	if err != nil {
		return models.DirectOrgan{}, err
	}

	if directOrgan.Code != updateDirectOrgan.Code {
		return models.DirectOrgan{}, errors.New("code cannot be changed")
	}

	directOrgan, err = directOrganCase.repo.Update(updateDirectOrgan)
	if err != nil {
		return models.DirectOrgan{}, err
	}
	return directOrgan, err
}

func (directOrganCase *DirectOrganCase) Delete(id uuid.UUID) error {
	var err error

	_, err = directOrganCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = directOrganCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
