package models

type Currency struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	ShortName        string
	Scale            int64
	ScaleName        string
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
