package session_routes

import (
	"goshort/handlers/session_handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	sessionGroup := router.Group("/api/session")
	{
		sessionGroup.POST("/register", session_handler.Register)
		sessionGroup.POST("/login", session_handler.Login)
	}
}
