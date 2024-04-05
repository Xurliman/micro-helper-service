package models

type EducationType struct {
	Model

	Code      string `gorm:"unique;not null"`
	Name      string
	FlexFinId string
}
