package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func (repo *AccountRepo) Create(account models.Account) (models.Account, error) {
	err := repo.db.Create(&account).Error
	return account, err
}

func (repo *AccountRepo) Get(id uuid.UUID) (models.Account, error) {
	var account models.Account
	err := repo.db.Where("id = ?", id).First(&account).Error

	return account, err
}

func (repo *AccountRepo) Update(account models.Account) (models.Account, error) {
	var dbAccount models.Account
	if err := repo.db.Where("id = ?", account.ID).First(&account).Error; err != nil {
		return dbAccount, err
	}

	dbAccount.Code = account.Code
	dbAccount.Name = account.Name
	dbAccount.Description = account.Description
	dbAccount.FlexFinId = account.FlexFinId

	err := repo.db.Save(dbAccount).Error
	return dbAccount, err
}

func (repo *AccountRepo) GetByCode(code string) (models.Account, error) {
	var account models.Account
	err := repo.db.Where("code = ?", code).First(&account).Error

	return account, err
}

func (repo *AccountRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.Account{}).Error
	return err
}
