package usecase

import (
	"errors"

	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegionCase struct {
	repo interfaces.RegionRepoInterface
}

func NewRegion(repo interfaces.RegionRepoInterface) interfaces.RegionCaseInterface {
	return &RegionCase{repo: repo}
}

func (regionCase RegionCase) Create(region models.Region) (models.Region, error) {
	if _, err := regionCase.repo.GetByCode(region.Code); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Region{}, errors.New("the code has already been taken")
	}

	return regionCase.repo.Create(region)
}

func (regionCase RegionCase) Get(id uuid.UUID) (models.Region, error) {
	var region models.Region
	var err error

	region, err = regionCase.repo.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Region{}, errors.New("no such region with the id supplied")
		}

		return models.Region{}, err
	}
	return region, nil
}

func (regionCase RegionCase) Update(updateRegion models.Region) (models.Region, error) {
	var region models.Region
	var err error

	region, err = regionCase.repo.Get(updateRegion.ID)
	if err != nil {
		return models.Region{}, err
	}

	if region.Code != updateRegion.Code {
		return models.Region{}, errors.New("code cannot be changed")
	}

	region, err = regionCase.repo.Update(updateRegion)
	if err != nil {
		return models.Region{}, err
	}
	return region, err
}

func (regionCase RegionCase) Delete(id uuid.UUID) error {
	var err error

	_, err = regionCase.repo.Get(id)
	if err != nil {
		return err
	}

	err = regionCase.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
