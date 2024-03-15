package models

type DirectOrgan struct {
	gorm.Model

	Code      int64
	Name      string
	ShortName string
	CrudDates string
	CBUCode   int64
}
