package user

type UserRepository interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
}
