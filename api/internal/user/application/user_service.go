package application

import (
	"github.com/ebubekir/go-talk/api/internal/user/domain"
)

type UserService struct {
	repo        domain.UserRepository
	userService *domain.UserService
}

func NewUserService(repo domain.UserRepository) *UserService {
	userService := &domain.UserService{}
	return &UserService{repo: repo, userService: userService}
}

func (us *UserService) CreateUser(firebaseId, name, email string) error {
	if user, err := us.userService.Create(firebaseId, name, email); err != nil {
		return err
	} else {
		return us.repo.Create(user)
	}
}

func (us *UserService) GetUserByEmail(email string) (*domain.User, error) {
	return us.repo.GetUserByEmail(email)
}
