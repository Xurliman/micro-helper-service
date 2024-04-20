package repository

import (
	"github.com/Xurliman/banking-microservice/internal/models"
	"github.com/Xurliman/banking-microservice/pkg/v1/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankBranchRepo struct {
	db *gorm.DB
}

func NewBankBranch(db *gorm.DB) interfaces.BankBranchRepoInterface {
	return &BankBranchRepo{db: db}
}

func (repo *BankBranchRepo) Create(bankBranch models.BankBranch) (models.BankBranch, error) {
	err := repo.db.Create(&bankBranch).Error
	return bankBranch, err
}

func (repo *BankBranchRepo) Get(id uuid.UUID) (models.BankBranch, error) {
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
	dbBankBranch.BankId = bankBranch.BankId
	dbBankBranch.RegionId = bankBranch.RegionId
	dbBankBranch.DistrictId = bankBranch.DistrictId
	dbBankBranch.Address = bankBranch.Address
	dbBankBranch.OpenDate = bankBranch.OpenDate
	dbBankBranch.CloseDate = bankBranch.CloseDate
	dbBankBranch.ActivationDate = bankBranch.ActivationDate
	dbBankBranch.DeactivationDate = bankBranch.DeactivationDate
	dbBankBranch.FlexFinId = bankBranch.FlexFinId

	err := repo.db.Save(dbBankBranch).Error
	return dbBankBranch, err
}

func (repo *BankBranchRepo) Delete(id uuid.UUID) error {
	err := repo.db.Where("id = ?", id).Delete(&models.BankBranch{}).Error
	return err
}

func (repo *BankBranchRepo) GetByCode(code string) (models.BankBranch, error) {
	var bankBranch models.BankBranch
	err := repo.db.Where("code = ?", code).First(&bankBranch).Error

	return bankBranch, err
}
