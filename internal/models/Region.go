package models

type Region struct {
	Model

	Code             int64 `gorm:"unique;not null"`
	Name             string
	CountryId        int64
	ActivationDate   string
	DeactivationDate string
	OldCode          string
	OldName          string
	FlexFinId        string
}
