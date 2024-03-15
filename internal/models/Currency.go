package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model

	Code      int64
	Name      string
	ShortName string
	Scale     int64
	ScaleName string
}
