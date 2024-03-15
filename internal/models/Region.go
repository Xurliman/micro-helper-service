package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model

	Code             int64
	Name             string
	CountryID        int64
	ActivationDate   string
	DeactivationDate string
	OldName          string
}
