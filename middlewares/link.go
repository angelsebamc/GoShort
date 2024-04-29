package middlewares

import (
	"goshort/dtos/link_dto"
	"goshort/services/link_service"
	"goshort/utils"
	"goshort/utils/json_response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		link_body, exists := c.Get("link_body")

		if !exists {
			c.JSON(http.StatusNotFound, json_response.New(http.StatusNotFound, "url not given", nil))
			return
		}

		link, _ := link_body.(link_dto.LinkDTO_Post)

		log.Printf(link.OriginalUrl)

		is_valid_url := utils.ValidateURL(link.OriginalUrl)

		if !is_valid_url {
			c.JSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "Invalid URL. Please provide de complete path", nil))
			return
		}
		c.Next()
	}
}

func UrlExistsForTheUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, exists := c.Get("user_id")

		if !exists {
			c.AbortWithStatusJSON(http.StatusNotFound, json_response.New(http.StatusNotFound, "user does not exists", nil))
			return
		}

		var link_from_body link_dto.LinkDTO_Post

		if err := c.BindJSON(&link_from_body); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		link, status := link_service.GetInstance().GetLinkByOriginalUrl(link_from_body.OriginalUrl)

		if status.Code == http.StatusOK && user_id == link.UserID {
			c.AbortWithStatusJSON(http.StatusConflict, json_response.New(http.StatusConflict, "you already have a short url for that url", link))
			return
		}

		c.Set("link_body", link_from_body)

		c.Next()
	}
}
