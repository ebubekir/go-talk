package http

import (
	"github.com/ebubekir/go-talk/api/internal/user/application"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, userService *application.UserService) {
	handler := NewUserHandler(userService)

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", handler.CreateUser())
		userGroup.GET("/me", handler.GetCurrentUser())
	}
}
