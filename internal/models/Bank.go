package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model

	Code             int64
	Name             string
	ShortName        string
	CountryID        int64
	OpenDate         string
	CloseDate        string
	ActivationDate   string
	DeactivationDate string
	CreatedAt        string
	UpdatedAt        string
}
