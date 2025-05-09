package postgres

import (
	"database/sql"

	domain "github.com/Zaharazov/AdvancedArchitectureTemplate/internal/domain/user"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return r.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepository) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
