package models

type Country struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	ShortName        string
	CurrencyId       int64
	CodeAlpha2       string
	CodeAlpha3       string
	TerritoryCode    int64
	ActivationDate   string
	DeactivationDate string
	FlexFinId        string
}
