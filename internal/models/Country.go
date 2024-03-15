package models

type Country struct {
	gorm.Model

	Code          int64
	Name          string
	ShortName     string
	CurrencyID    int64
	CodeAlpha2    string
	CodeAlpha3    string
	TerritoryCode int64
	CrudDates     string
}
