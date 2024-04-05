package models

type BankBranch struct {
	Model

	Code             int64 `gorm:"unique;not null"`
	Name             string
	BankId           int64
	RegionId         int64
	DistrictId       int64
	Address          string
	OpenDate         string
	CloseDate        string
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
