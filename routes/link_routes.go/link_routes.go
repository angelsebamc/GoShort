package link_routes

import (
	"goshort/handlers/link_handler"
	"goshort/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	user_group := router.Group("/api/link")
	{
		user_group.POST("/create", middlewares.ValidateURL(), middlewares.Auth(), middlewares.UrlExistsForTheUser(), link_handler.CreateLink)
		user_group.DELETE("/delete", middlewares.Auth(), link_handler.DeleteLink)
		user_group.GET("/user_links", middlewares.Auth(), link_handler.GetLinksByUserId)
	}

	//shorten url
	router.GET("/:short_url", link_handler.ShortUrlRedirect)
}
