package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `gorm:"primaryKey;type:binary(16);default:(UUID_TO_BIN(UUID(),TRUE));not null;"`
	CreatedAt string
	UpdatedAt string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
