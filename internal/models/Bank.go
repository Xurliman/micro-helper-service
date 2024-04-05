package models

type Bank struct {
	Model

	Code             int64 `gorm:"unique;not null"`
	Name             string
	ShortName        string
	CountryID        int64
	OpenDate         string
	CloseDate        string
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
