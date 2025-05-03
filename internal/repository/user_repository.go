package repository

import "github.com/Zaharazov/AdvancedArchitectureTemplate/internal/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int64) (*model.User, error)
}
