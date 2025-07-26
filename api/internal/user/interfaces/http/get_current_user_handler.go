package http

import (
	"github.com/ebubekir/go-talk/api/internal/httpctx"
	"github.com/ebubekir/go-talk/api/internal/response"
	"github.com/gin-gonic/gin"
)

type GetCurrentUserRequest struct{}

// GetCurrentUser
// @Summary      GetCurrentUser
// @Description  Get  current authenticated user.
// @Security 	 ApiKeyAuth
// @Tags         User
// @Success      200     {object} UserResponse
// @Failure      default {object} response.ApiError
// @Router       /user/me [get]
func (us *UserHandler) GetCurrentUser() gin.HandlerFunc {
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
