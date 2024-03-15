package models

type NationalEconomySector struct {
	gorm.Model

	Code             int64
	Name             string
	CBUCode          int64
	CBUGroupCode     int64
	ActivationDate   string
	DeactivationDate string
	CBUReferenceKey  int
}
