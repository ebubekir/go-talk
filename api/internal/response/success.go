package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusResponse struct {
	Message string `json:"message" validate:"required"` //Message
	Code    int    `json:"code" validate:"required"`    //Code of status which is always 200
} //@name StatusResponse

func Status(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &StatusResponse{
		Message: msg,
		Code:    http.StatusOK,
	})
}

func Success(c *gin.Context, response any) {
	c.JSON(http.StatusOK, response)
}
