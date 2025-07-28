package domain

type UserRepository interface {
	Create(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
}
