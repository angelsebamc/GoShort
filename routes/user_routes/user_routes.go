package user_routes

import (
	"goshort/handlers/user_handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	sessionGroup := router.Group("/api/user")
	{
		sessionGroup.POST("/signin", user_handler.SignIn)
		sessionGroup.POST("/auth", user_handler.Auth)
	}
}
