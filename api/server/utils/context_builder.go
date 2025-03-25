package utils

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func BuildInternalContext(ctx *gin.Context) context.Context{
		// Создаем контекст с таймаутом 2 секунды
		internalCtx, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)

		go func() {
			<-ctx.Done() 
			cancel()   
		}()
	
		return internalCtx
}

