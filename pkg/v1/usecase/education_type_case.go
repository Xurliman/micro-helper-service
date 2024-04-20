package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EducationTypeCase struct {
	repo interfaces.EducationTypeRepoInterface
}

func NewEducationType(repo interfaces.EducationTypeRepoInterface) interfaces.EducationTypeCaseInterface {
	return &EducationTypeCase{repo: repo}
}

func (educationTypeCase EducationTypeCase) Create(educationType models.EducationType) (models.EducationType, error) {
	if _, err := educationTypeCase.repo.GetByCode(educationType.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.EducationType{}, errors.New("the code has already been taken")
	}

	return educationTypeCase.repo.Create(educationType)
}

func (educationTypeCase EducationTypeCase) Get(id uuid.UUID) (models.EducationType, error) {
	var educationType models.EducationType
	var err error

	educationType, err = educationTypeCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.EducationType{}, errors.New("no such educationType with the id supplied")
		}

		return models.EducationType{}, err
	}
	return educationType, nil
}

func (educationTypeCase EducationTypeCase) Update(updateEducationType models.EducationType) (models.EducationType, error) {
	var educationType models.EducationType
	var err error

	educationType, err = educationTypeCase.repo.Get(updateEducationType.ID)
	if err != nil {
		return models.EducationType{}, err
	}

	if educationType.Code != updateEducationType.Code {
		return models.EducationType{}, errors.New("code cannot be changed")
	}

	educationType, err = educationTypeCase.repo.Update(updateEducationType)
	if err != nil {
		return models.EducationType{}, err
	}
	return educationType, err
}

func (educationTypeCase EducationTypeCase) Delete(id uuid.UUID) error {
	var err error

	_, err = educationTypeCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = educationTypeCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
