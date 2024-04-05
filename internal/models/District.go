package models

type District struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	RegionId         int64
	OldCode          int64
	OldName          string
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
