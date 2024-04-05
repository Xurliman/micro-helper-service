package models

type PassportType struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	OldCode          string
	OldName          string
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
