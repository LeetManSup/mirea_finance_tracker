package model

import (
	"time"

	"github.com/google/uuid"
)

type RecurringTransaction struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID  `gorm:"type:uuid;not null"`
	AccountID      uuid.UUID  `gorm:"type:uuid;not null"`
	CategoryID     *uuid.UUID `gorm:"type:uuid"`
	Amount         float64    `gorm:"not null"`
	CurrencyCode   string     `gorm:"type:char(3);not null"`
	Description    string
	Interval       string    `gorm:"type:text;check:interval IN ('daily','weekly','monthly')"`
	StartDate      time.Time `gorm:"not null"`
	EndDate        *time.Time
	NextOccurrence time.Time `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`

	Account  Account
	Category *Category
	Currency Currency
}
