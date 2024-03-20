package interfaces

import "github.com/Xurliman/banking-microservice/internal/models"

type ClientTypeClassifierRepoInterface interface {
	Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	Get(id int64) (models.ClientTypeClassifier, error)
	Update(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	GetByCode(code string) (models.ClientTypeClassifier, error)
	Delete(id int64) error
}

type ClientTypeClassifierCaseInterface interface {
	Create(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	Get(id int64) (models.ClientTypeClassifier, error)
	Update(classifier models.ClientTypeClassifier) (models.ClientTypeClassifier, error)
	Delete(id int64) error
}
