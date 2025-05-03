package service

import (
	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/domain/model"
	"github.com/Zaharazov/AdvancedArchitectureTemplate/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetUserByID(id int64) (*model.User, error) {
	return s.Repo.GetByID(id)
}
