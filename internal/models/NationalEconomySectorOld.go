package models

type NationalEconomySectorOld struct {
	Model

	Code             int64 `gorm:"unique;not null"`
	Name             string
	CBUCode          int64
	CBUGroupCode     int64
	ActivationDate   string
	DeactivationDate string
	CBUReferenceKey  int32
	FlexFinId        string
}
