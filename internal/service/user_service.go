package service

import (
	"mirea_finance_tracker/internal/model"
	"mirea_finance_tracker/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) GetByID(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}
