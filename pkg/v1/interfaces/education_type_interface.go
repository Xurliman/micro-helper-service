package interfaces

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
)

type EducationTypeRepoInterface interface {
	Create(educationType models.EducationType) (models.EducationType, error)
	Get(id uuid.UUID) (models.EducationType, error)
	Update(educationType models.EducationType) (models.EducationType, error)
	GetByCode(code string) (models.EducationType, error)
	Delete(id uuid.UUID) error
}

type EducationTypeCaseInterface interface {
	Create(educationType models.EducationType) (models.EducationType, error)
	Get(id uuid.UUID) (models.EducationType, error)
	Update(educationType models.EducationType) (models.EducationType, error)
	Delete(id uuid.UUID) error
}
