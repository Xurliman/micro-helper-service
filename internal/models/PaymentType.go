package models

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model

	Code             int64
	Name             string
	ActivationDate   string
	DeactivationDate string
	OldCode          int64
	OldName          string
	//NameUz          string
}
