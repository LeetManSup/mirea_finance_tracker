package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	Type      string    `gorm:"type:text;check:type IN ('income','expense')"`
	IsDefault bool
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Transactions []Transaction `gorm:"foreignKey:CategoryID"`
}
