package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt string
	UpdatedAt string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
