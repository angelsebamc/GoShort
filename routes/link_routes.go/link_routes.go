package link_routes

import (
	"goshort/handlers/link_handler"
	"goshort/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	sessionGroup := router.Group("/api/link")
	{
		sessionGroup.POST("/create", middlewares.Auth(), middlewares.UrlExistsForTheUser(), link_handler.CreateNewShortUrl)
	}
}
