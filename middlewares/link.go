package middlewares

import (
	"goshort/dtos/link_dto"
	"goshort/services/link_service"
	"goshort/utils/json_response"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ValidateURL() gin.HandlerFunc {
	return func(c *gin.Context) {
		var link_from_body link_dto.LinkDTO_Post

		if err := c.BindJSON(&link_from_body); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "No URL provided", nil))
			return
		}

		if _, err := url.ParseRequestURI(link_from_body.OriginalUrl); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "invalid URL. Please provide the complete path", nil))
			return
		}

		c.Set("link_from_body", link_from_body)

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

		link_from_body, exists := c.Get("link_from_body")

		if !exists {
			c.JSON(http.StatusInternalServerError, json_response.New(http.StatusNotFound, "url not given", nil))
		}

		link_parsed, _ := link_from_body.(link_dto.LinkDTO_Post)

		link, status := link_service.GetInstance().GetLinkByOriginalUrl(link_parsed.OriginalUrl)

		if status.Code == http.StatusOK && user_id == link.UserID {
			c.AbortWithStatusJSON(http.StatusConflict, json_response.New(http.StatusConflict, "you already have a short url for that url", link))
			return
		}

		c.Set("link_body", link_from_body)

		c.Next()
	}
}
