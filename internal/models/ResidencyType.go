package models

type ResidencyType struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
	NameUz           string
}