package models

type PaymentType struct {
	Model

	Code             int64 `gorm:"unique;not null"`
	Name             string
	ActivationDate   string
	DeactivationDate string
	OldCode          int64
	OldName          string
	NameUz           string
	FlexFinId        string
}
