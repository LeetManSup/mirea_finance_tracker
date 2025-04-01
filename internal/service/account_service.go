package service

import (
	"errors"

	"mirea_finance_tracker/internal/model"
	"mirea_finance_tracker/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AccountService struct {
	accountRepo  *repository.AccountRepository
	currencyRepo *repository.CurrencyRepository
}

func NewAccountService(accountRepo *repository.AccountRepository, currencyRepo *repository.CurrencyRepository) *AccountService {
	return &AccountService{accountRepo, currencyRepo}
}

type UpdateAccountInput struct {
	Name           *string
	CurrencyCode   *string
	InitialBalance *float64
}

func (s *AccountService) CreateAccount(userID, name, currencyCode string, initialBalance float64) (uuid.UUID, error) {
	exists, err := s.currencyRepo.Exists(currencyCode)
	if err != nil || !exists {
		logrus.Errorf("Ошибка, неизвестная валюта: %v", err)
		return uuid.Nil, errors.New("invalid currency code")
	}

	account := model.Account{
		ID:             uuid.New(),
		UserID:         uuid.MustParse(userID),
		Name:           name,
		CurrencyCode:   currencyCode,
		InitialBalance: initialBalance,
	}

	logrus.Infof("Создание нового счёта: %s", account.Name)

	err = s.accountRepo.Create(&account)
	if err != nil {
		logrus.Errorf("Ошибка при создании счёта: %v", err)
		return uuid.Nil, err
	}

	logrus.Infof("Счёт успешно создан: %s", account.Name)
	return account.ID, nil
}

func (s *AccountService) GetAccountsByUser(userID string) ([]model.Account, error) {
	return s.accountRepo.GetByUserID(userID)
}

func (s *AccountService) GetAccountByID(userID, accountID string) (*model.Account, error) {
	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return nil, err
	}

	if account.UserID.String() != userID {
		return nil, errors.New("access denied")
	}

	return account, nil
}

func (s *AccountService) DeleteAccount(userID, accountID string) error {
	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return err
	}

	if account.UserID.String() != userID {
		return errors.New("access denied")
	}

	return s.accountRepo.Delete(accountID)
}

func (s *AccountService) UpdateAccount(userID, accountID string, input UpdateAccountInput) error {
	account, err := s.accountRepo.GetByID(accountID)
	if err != nil {
		return err
	}

	if account.UserID.String() != userID {
		return errors.New("access denied")
	}

	if input.CurrencyCode != nil {
		ok, err := s.currencyRepo.Exists(*input.CurrencyCode)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("invalid currency code")
		}
		account.CurrencyCode = *input.CurrencyCode
	}

	if input.Name != nil {
		account.Name = *input.Name
	}

	if input.InitialBalance != nil {
		account.InitialBalance = *input.InitialBalance
	}

	return s.accountRepo.Update(account)
}
