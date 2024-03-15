package models

type ClientTypeClassifier struct {
	gorm.Model

	Code             int64
	Name             string
	ShortName        string
	ClientType       int
	ActivationDate   string
	DeactivationDate string
	CBUReferenceKey  int
	OldCode          int
	OldName          int
	//NameUz           int
}
