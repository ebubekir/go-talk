package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required"`
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

		response.Success(c, &UserResponse{
			Name:  requestUser.Name,
			Id:    requestUser.Id,
			Email: requestUser.Email,
		})
	}
}
