package link_handler

import (
	"goshort/dtos/link_dto"
	"goshort/services/link_service"
	"goshort/utils/json_response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateNewShortUrl(c *gin.Context) {
	link_body, _ := c.Get("link_body")

	original_link := link_body.(link_dto.LinkDTO_Post)

	new_link, status := link_service.GetInstance().CreateShortURL(c, &original_link)

	if status.Code != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, json_response.New(status.Code, status.Message, nil))
	}

	c.JSON(http.StatusCreated, json_response.New(status.Code, status.Message, new_link))
}

func ShortUrlRedirect(c *gin.Context) {
	short_url_param := c.Param("short_url")

	join_with_base := os.Getenv("BASE_URL") + "/" + short_url_param

	link, status := link_service.GetInstance().GetLinkByShortUrl(join_with_base)

	if status.Code != http.StatusOK {
		c.JSON(http.StatusInternalServerError, json_response.New(status.Code, status.Message, nil))
	}

	c.Redirect(http.StatusPermanentRedirect, link.OriginalUrl)
}
