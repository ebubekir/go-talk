package http

import "github.com/ebubekir/go-talk/api/internal/user/application"

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
