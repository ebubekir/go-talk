package middleware

import (
	"context"
	"errors"
	keys "github.com/ebubekir/go-talk/api/internal/keys"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/ebubekir/go-talk/api/internal/user/application"
	"github.com/ebubekir/go-talk/api/pkg/firebase"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthMiddleware struct {
	firebaseApp *firebase.App
	userService *application.UserService
}

func NewAuthMiddleware(firebaseApp *firebase.App, userService *application.UserService) *AuthMiddleware {
	return &AuthMiddleware{firebaseApp: firebaseApp, userService: userService}
}

func (a *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		if authToken == "" {
			// No Auth token
			response.UnauthorizedError(c, errors.New("authorization token is required"))
			return
		}

		idToken := strings.TrimSpace(strings.Replace(authToken, "Bearer", "", 1))
		if idToken == "" {
			response.UnauthorizedError(c, errors.New("Authorization token is required"))
			return
		}

		firebaseToken, err := a.firebaseApp.AuthClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			response.UnauthorizedError(c, err)
			return
		}

		name := ""
		email := ""

		user, err := a.userService.GetUserByEmail(email)
		if err != nil {
			response.SystemError(c, err)
			return
		}

		if user == nil {
			// Google login add user to database
			if v, hasValue := firebaseToken.Claims["name"]; hasValue {
				name = v.(string)
				c.Set(keys.Name, name)
			}

			err = a.userService.CreateUser(
				firebaseToken.UID,
				name,
				email,
			)

		}

		if user.IsDeleted {
			response.UnauthorizedError(c, errors.New("user is deleted"))
			return
		}

		c.Set(keys.User, user)
		c.Next()
	}
}
