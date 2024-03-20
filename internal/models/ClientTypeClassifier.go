package models

import "gorm.io/gorm"

type ClientTypeClassifier struct {
	gorm.Model

	Code             int64
	Name             string
	ShortName        string
	ClientType       int32
	ActivationDate   string
	DeactivationDate string
	CBUReferenceKey  int32
	OldCode          int32
	OldName          string
	//NameUz           int
}
