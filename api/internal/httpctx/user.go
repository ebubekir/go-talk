package httpctx

import (
	"github.com/ebubekir/go-talk/api/internal/commons"
	"github.com/ebubekir/go-talk/api/internal/user/domain"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) *domain.User {
	return c.MustGet(commons.UserContextKey).(*domain.User)
}
