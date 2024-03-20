package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"gorm.io/gorm"
)

type BankBranchRepo struct {
	db *gorm.DB
}

func (repo *BankBranchRepo) Create(bankBranch models.BankBranch) (models.BankBranch, error) {
	err := repo.db.Create(&bankBranch).Error
	return bankBranch, err
}

func (repo *BankBranchRepo) Get(id int64) (models.BankBranch, error) {
	var bankBranch models.BankBranch
	err := repo.db.Where("id = ?", id).First(&bankBranch).Error

	return bankBranch, err
}

func (repo *BankBranchRepo) Update(bankBranch models.BankBranch) (models.BankBranch, error) {
	var dbBankBranch models.BankBranch
	if err := repo.db.Where("id = ?", bankBranch.ID).First(&bankBranch).Error; err != nil {
		return dbBankBranch, err
	}

	dbBankBranch.Code = bankBranch.Code
	dbBankBranch.Name = bankBranch.Name
	dbBankBranch.BankID = bankBranch.BankID
	dbBankBranch.RegionID = bankBranch.RegionID
	dbBankBranch.DistrictID = bankBranch.DistrictID
	dbBankBranch.Address = bankBranch.Address
	dbBankBranch.OpenDate = bankBranch.OpenDate
	dbBankBranch.CloseDate = bankBranch.CloseDate
	dbBankBranch.CrudDates = bankBranch.CrudDates

	err := repo.db.Save(dbBankBranch).Error
	return dbBankBranch, err
}

func (repo *BankBranchRepo) Delete(id int64) error {
	err := repo.db.Where("id = ?", id).Delete(&models.BankBranch{}).Error
	return err
}

func (repo *BankBranchRepo) GetByCode(code string) (models.BankBranch, error) {
	var bankBranch models.BankBranch
	err := repo.db.Where("code = ?", code).First(&bankBranch).Error

	return bankBranch, err
}
