package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
} // @name LoginRequest

type UserResponse struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"e164"`
} // @name User

// CreateUser
// @Summary      CreateUser
// @Description  Create user.
// @Security 	 ApiKeyAuth
// @Tags         User
// @Success      200     {object} UserResponse
// @Failure      default {object} response.ApiError
// @Router       /user/create [post]
func (us *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUser := httpctx.GetUserFromContext(c)

		if requestUser == nil {
			response.BadRequestWithMessage(c, "user not found")
			return
		}
	}
}
