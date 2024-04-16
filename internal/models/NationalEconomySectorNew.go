package models

type NationalEconomySectorNew struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	NameUz           string
	Section          int32
	SectionCode      string
	SectionName      string
	SectionNameUz    string
	GroupCode        string
	GroupName        string
	GroupNameUz      string
	Group1Code       string
	Group1Name       string
	Group1NameUz     string
	Group2Code       string
	Group2Name       string
	Group2NameUz     string
	ActivationDate   string
	DeactivationDate string
	CBUReferenceKey  int32
	FlexFinId        string
}
