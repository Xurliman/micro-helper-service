package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientTypeClassifierCase struct {
	repo interfaces.ClientTypeClassifierRepoInterface
}

func NewClientTypeClassifier(repo interfaces.ClientTypeClassifierRepoInterface) interfaces.ClientTypeClassifierCaseInterface {
	return &ClientTypeClassifierCase{repo: repo}
}

func (classifierCase *ClientTypeClassifierCase) Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error) {
	if _, err := classifierCase.repo.GetByCode(classifier.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ClientTypeClassifier{}, errors.New("the code has already been taken")
	}

	return classifierCase.repo.Create(classifier)
}

func (classifierCase *ClientTypeClassifierCase) Get(id uuid.UUID) (models.ClientTypeClassifier, error) {
	var classifier models.ClientTypeClassifier
	var err error

	classifier, err = classifierCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ClientTypeClassifier{}, errors.New("no such classifier with the id supplied")
		}

		return models.ClientTypeClassifier{}, err
	}
	return classifier, nil
}

func (classifierCase *ClientTypeClassifierCase) Update(updateClassifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error) {
	var classifier models.ClientTypeClassifier
	var err error

	classifier, err = classifierCase.repo.Get(updateClassifier.ID)
	if err != nil {
		return models.ClientTypeClassifier{}, err
	}

	if classifier.Code != updateClassifier.Code {
		return models.ClientTypeClassifier{}, errors.New("code cannot be changed")
	}

	classifier, err = classifierCase.repo.Update(updateClassifier)
	if err != nil {
		return models.ClientTypeClassifier{}, err
	}
	return classifier, err
}

func (classifierCase *ClientTypeClassifierCase) Delete(id uuid.UUID) error {
	var err error

	_, err = classifierCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = classifierCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
