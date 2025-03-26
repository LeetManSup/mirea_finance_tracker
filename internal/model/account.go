package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	Name           string    `gorm:"not null"`
	CurrencyCode   string    `gorm:"type:char(3);not null"`
	InitialBalance float64
	CreatedAt      time.Time `gorm:"autoCreateTime"`

	Currency     Currency
	Transactions []Transaction `gorm:"foreignKey:AccountID"`
}
