package support

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
    supportContreoller *SupportContreoller
}

func NewRouter(supportContreoller *SupportContreoller) *Router {
    return &Router{
        supportContreoller: supportContreoller,
    }
}

func (c *Router) RegisterRoutes(rg *gin.RouterGroup) {
    rg.GET("/ai-assist", c.supportContreoller.ProcessQuestion)
}