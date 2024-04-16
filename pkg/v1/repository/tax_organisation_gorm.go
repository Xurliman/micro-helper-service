package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaxOrganisationRepo struct {
	db *gorm.DB
}

func (repo *TaxOrganisationRepo) Create(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error) {
	err := repo.db.Create(&taxOrganisation).Error
	return taxOrganisation, err
}

func (repo *TaxOrganisationRepo) Get(id uuid.UUID) (models.TaxOrganisation, error) {
	var taxOrganisation models.TaxOrganisation
	err := repo.db.Where("id = ?", id).First(&taxOrganisation).Error

	return taxOrganisation, err
}

func (repo *TaxOrganisationRepo) Update(taxOrganisation models.TaxOrganisation) (models.TaxOrganisation, error) {
	var dbTaxOrganisation models.TaxOrganisation
	if err := repo.db.Where("id = ?", taxOrganisation.ID).First(&taxOrganisation).Error; err != nil {
		return dbTaxOrganisation, err
	}

	dbTaxOrganisation.Code = taxOrganisation.Code
	dbTaxOrganisation.Name = taxOrganisation.Name
	dbTaxOrganisation.FlexFinId = taxOrganisation.FlexFinId

	err := repo.db.Save(dbTaxOrganisation).Error
	return dbTaxOrganisation, err
}

func (repo *TaxOrganisationRepo) GetByCode(code string) (models.TaxOrganisation, error) {
	var taxOrganisation models.TaxOrganisation
	err := repo.db.Where("code = ?", code).First(&taxOrganisation).Error

	return taxOrganisation, err
}

func (repo *TaxOrganisationRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.TaxOrganisation{}).Error
	return err
}
