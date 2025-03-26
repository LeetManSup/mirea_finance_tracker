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

func (r *AccountRepository) GetByID(id string) (*model.Account, error) {
	var account model.Account
	err := r.db.First(&account, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepository) Delete(id string) error {
	return r.db.Delete(&model.Account{}, "id = ?", id).Error
}

func (r *AccountRepository) Update(account *model.Account) error {
	return r.db.Save(account).Error
}
