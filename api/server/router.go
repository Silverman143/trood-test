package server

import (
	"log/slog"
	"time"
	"trood-test/api/server/controllers/support"
	"trood-test/api/server/response"
	"trood-test/internal/services/nlp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



type Handler struct{
	log *slog.Logger
	nlpService *nlp.NLPService
}

func NewHandler( log *slog.Logger, nlpService *nlp.NLPService) *Handler {
	return &Handler{
		log: log,
		nlpService: nlpService,
	}
}

func (c *Handler) InitRouts() *gin.Engine{
	router := gin.New()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	   }))

	router.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "43200")
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// health checker
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := router.Group("/v1")


	responseBuilder := response.NewResponseBuilder()

	supportController := support.NewController(c.log, responseBuilder, c.nlpService)
	supportRouter := support.NewRouter(supportController)
	supportRouter.RegisterRoutes(v1)

	return router
}