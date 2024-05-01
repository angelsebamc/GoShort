package user_routes

import (
	"goshort/handlers/user_handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	user_group := router.Group("/api/user")
	{
		user_group.POST("/signin", user_handler.SignIn)
		user_group.POST("/auth", user_handler.Auth)
	}
}
