package link_routes

import (
	"goshort/handlers/link_handler"
	"goshort/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	sessionGroup := router.Group("/api/link")
	{
		sessionGroup.POST("/create", middlewares.ValidateURL(), middlewares.Auth(), middlewares.UrlExistsForTheUser(), link_handler.CreateLink)
		sessionGroup.DELETE("/delete", middlewares.Auth(), link_handler.DeleteLink)
		sessionGroup.GET("/user_links", middlewares.Auth(), link_handler.GetLinksByUserId)
	}

	//shorten url
	router.GET("/:short_url", link_handler.ShortUrlRedirect)
}
