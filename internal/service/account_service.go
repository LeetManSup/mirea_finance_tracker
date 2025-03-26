package service

import (
	"errors"
	"mirea_finance_tracker/internal/model"
	"mirea_finance_tracker/internal/repository"

	"github.com/google/uuid"
)

type AccountService struct {
	accountRepo  *repository.AccountRepository
	currencyRepo *repository.CurrencyRepository
}

func NewAccountService(accountRepo *repository.AccountRepository, currencyRepo *repository.CurrencyRepository) *AccountService {
	return &AccountService{accountRepo, currencyRepo}
}

func (s *AccountService) CreateAccount(userID, name, currencyCode string, initialBalance float64) (uuid.UUID, error) {
	// Проверка валюты
	exists, err := s.currencyRepo.Exists(currencyCode)
	if err != nil || !exists {
		return uuid.Nil, errors.New("invalid currency code")
	}

	account := model.Account{
		ID:             uuid.New(),
		UserID:         uuid.MustParse(userID),
		Name:           name,
		CurrencyCode:   currencyCode,
		InitialBalance: initialBalance,
	}

	err = s.accountRepo.Create(&account)
	if err != nil {
		return uuid.Nil, err
	}

	return account.ID, nil
}

func (s *AccountService) GetAccountsByUser(userID string) ([]model.Account, error) {
	return s.accountRepo.GetByUserID(userID)
}
