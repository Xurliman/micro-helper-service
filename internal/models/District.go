package models

import "gorm.io/gorm"

type District struct {
	gorm.Model

	Code             int64
	Name             string
	RegionID         int64
	ActivationDate   string
	DeactivationDate string
	OldCode          int64
}
