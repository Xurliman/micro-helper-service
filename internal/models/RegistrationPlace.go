package models

type RegistrationPlace struct {
	Model

	Code      string `gorm:"unique;not null"`
	Name      string
	FlexFinId string
}
