package middlewares

import (
	"goshort/dtos/link_dto"
	"goshort/services/link_service"
	"goshort/utils/json_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlExistsForTheUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, exists := c.Get("user_id")

		if !exists {
			c.AbortWithStatusJSON(403, json_response.New(404, "user does not exists", nil))
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
