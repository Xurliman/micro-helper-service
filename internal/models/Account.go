package models

type Account struct {
	Model

	Code        string `json:"code" validate:"required" gorm:"unique;not null"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	FlexFinId   string `json:"flex_fin_id" validate:"required"`
}
