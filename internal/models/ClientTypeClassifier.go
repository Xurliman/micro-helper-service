package models

type ClientTypeClassifier struct {
	Model

	Code             string `gorm:"unique;not null"`
	Name             string
	ShortName        string
	ClientType       int32
	ActivationDate   string
	DeactivationDate string
	CBUReferenceKey  int32
	OldCode          int32
	OldName          string
	NameUz           string
	FlexFinId        string
}
