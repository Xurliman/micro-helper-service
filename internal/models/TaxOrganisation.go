package models

type TaxOrganisation struct {
	Model

	Code      string `gorm:"unique;not null"`
	Name      string
	FlexFinId string
}
