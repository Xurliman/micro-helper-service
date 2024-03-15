package models

type BankBranch struct {
	gorm.Model

	Code       int64
	Name       string
	BankID     int64
	RegionID   int64
	DistrictID int64
	Address    string
	OpenDate   string
	CloseDate  string
	CrudDates  string
	//CreatedAt  string
	//UpdatedAt  string
}
