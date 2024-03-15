package models

import "gorm.io/gorm"

type ClientTypeClassifier struct {
	gorm.Model

	Code int64
	Name string
}
