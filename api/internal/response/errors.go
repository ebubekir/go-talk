package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ApiError struct {
	Message string `json:"message" validate:"required"`
	Code    int    `json:"code" validate:"required"`
} //@name Error

func newApiError(statusCode int, err error) *ApiError {
	return newApiErrorFromMessage(statusCode, err.Error())
}

func newApiErrorFromMessage(statusCode int, errorMessage string) *ApiError {
	if len(errorMessage) != 0 {
		errorMessage = strings.ToUpper(errorMessage[:1]) + errorMessage[1:]
	}

	return &ApiError{
		Message: errorMessage,
		Code:    statusCode,
	}
}

func SystemError(c *gin.Context, err error) {
	_ = c.Error(err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, newApiError(http.StatusInternalServerError, err))
}

func BadRequest(c *gin.Context, err error) {
	ErrorWithStatusCodeAndMessage(c, http.StatusBadRequest, err.Error())
}

func BadRequestWithMessage(c *gin.Context, msg string) {
	ErrorWithStatusCodeAndMessage(c, http.StatusBadRequest, msg)
}

func ValidationError(c *gin.Context, err error) {
	ErrorWithStatusCodeAndMessage(c, http.StatusBadRequest, err.Error())
}

func UnauthorizedError(c *gin.Context, err error) {
	ErrorWithStatusCodeAndMessage(c, http.StatusUnauthorized, err.Error())
}

func ErrorWithStatusCodeAndMessage(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, newApiErrorFromMessage(statusCode, msg))
}
