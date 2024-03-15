package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model

	Code int64
	Name string
}
