package domain

import "github.com/google/uuid"

type UserService struct{}

func (us *UserService) Create(firebaseId, name, email string) (*User, error) {
	return &User{
		Id:         uuid.NewString(),
		FirebaseId: firebaseId,
		Name:       name,
		Email:      email,
		IsDeleted:  false,
	}, nil
}
