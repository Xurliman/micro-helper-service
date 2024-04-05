package models

type Account struct {
	Model
	Code        string `gorm:"unique;not null"`
	Name        string
	Description string
}
