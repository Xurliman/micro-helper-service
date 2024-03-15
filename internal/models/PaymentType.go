package models

type PaymentType struct {
	gorm.Model

	Code             int64
	Name             string
	ActivationDate   string
	DeactivationDate int64
	OldCode          string
	OldName          string
	//NameUz          string
}
