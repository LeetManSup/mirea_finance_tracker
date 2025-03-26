package repository

import (
	"mirea_finance_tracker/internal/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) Create(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) GetByUserID(userID string) ([]model.Account, error) {
	var accounts []model.Account
	err := r.db.Where("user_id = ?", userID).Find(&accounts).Error
	return accounts, err
}
