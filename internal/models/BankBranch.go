package models

import "gorm.io/gorm"

type BankBranch struct {
	gorm.Model

	Code int64
	Name string
}
