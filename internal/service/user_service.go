package service

import (
	domain "github.com/Zaharazov/AdvancedArchitectureTemplate/internal/domain/user"
)

type UserService struct {
	Repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetUserByID(id int64) (*domain.User, error) {
	return s.Repo.GetByID(id)
}
