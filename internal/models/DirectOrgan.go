package models

type DirectOrgan struct {
	Model

	Code             int64 `gorm:"unique;not null"`
	Name             string
	ShortName        string
	CrudDates        string
	CBUCode          int64
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
