package model

type Currency struct {
	Code   string `gorm:"primaryKey;type:char(3)"`
	Name   string `gorm:"not null"`
	Symbol string
}
