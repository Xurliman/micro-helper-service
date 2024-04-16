package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type ClientTypeClassifierRepoInterface interface {
	Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	Get(id uuid.UUID) (models.ClientTypeClassifier, error)
	Update(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	GetByCode(code string) (models.ClientTypeClassifier, error)
	Delete(id uuid.UUID) error
}

type ClientTypeClassifierCaseInterface interface {
	Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	Get(id uuid.UUID) (models.ClientTypeClassifier, error)
	Update(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	Delete(id uuid.UUID) error
}
