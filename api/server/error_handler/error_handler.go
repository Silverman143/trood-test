package error_handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type errorWithCodeResponse struct {
	Message   string `json:"message"`    // Человеко-читаемое сообщение
	ErrorCode string `json:"error_code"` // Уникальный код ошибки
}

func NewErrorResponseWithCode(ctx *gin.Context, statusCode int, message, errorCode string) {
	ctx.AbortWithStatusJSON(statusCode, errorWithCodeResponse{
		Message:   message,
		ErrorCode: errorCode,
	})
}

func NewErrorResponse(c *gin.Context, statusCode int, message string){
	
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}