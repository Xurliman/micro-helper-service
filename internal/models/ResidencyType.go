package models

type ResidencyType struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	ActivationDate   string
	DeactivationDate string
	NameUz           string
	FlexFinId        string
}
