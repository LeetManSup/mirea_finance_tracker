package repository

import (
	"gorm.io/gorm"
)

type CurrencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{db}
}

func (r *CurrencyRepository) Exists(code string) (bool, error) {
	var count int64
	err := r.db.Table("currencies").Where("code = ?", code).Count(&count).Error
	return count > 0, err
}
