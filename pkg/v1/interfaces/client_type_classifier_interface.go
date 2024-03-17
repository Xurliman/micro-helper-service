package interfaces

type ClientTypeClassifierInterface interface {
	Create(classifier models.ClientTypeClassifier) error
	Get(id string) error
	Update(classifier models.ClientTypeClassifier) error
	Delete(id string) error
}
