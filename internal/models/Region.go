package models

type Region struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	CountryId        int64
	ActivationDate   string
	DeactivationDate string
	OldCode          string
	OldName          string
	FlexFinId        string
}
