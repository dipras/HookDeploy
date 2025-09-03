package routes

import (
	"HookDeploy/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.POST("/webhook", handlers.WebhookHandler)
	}
}
