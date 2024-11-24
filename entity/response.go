package entity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Message: message,
		Data:    data,
	})
}

func RespondError(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, APIResponse{
		Message: message,
		Error:   err,
	})
}
