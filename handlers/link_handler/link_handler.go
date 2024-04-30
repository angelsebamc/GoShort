package link_handler

import (
	"goshort/dtos/link_dto"
	"goshort/services/link_service"
	"goshort/utils/json_response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateLink(c *gin.Context) {
	link_body, _ := c.Get("link_body")

	link := link_body.(link_dto.LinkDTO_Post)

	new_link, status := link_service.GetInstance().CreateLink(c, &link)

	if status.Code != http.StatusCreated {
		c.JSON(int(status.Code), json_response.New(status.Code, status.Message, nil))
	}

	c.JSON(int(status.Code), json_response.New(status.Code, status.Message, new_link))
}

func DeleteLink(c *gin.Context) {
	var link_from_body link_dto.LinkDTO_Delete

	if err := c.BindJSON(&link_from_body); err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "No id provided", nil))
		return
	}

	link_id := link_from_body.ID

	deleted_link, status := link_service.GetInstance().DeleteLinkById(link_id)

	if status.Code != http.StatusCreated {
		c.JSON(int(status.Code), json_response.New(status.Code, status.Message, nil))
	}

	c.JSON(int(status.Code), json_response.New(status.Code, status.Message, deleted_link))
}

func ShortUrlRedirect(c *gin.Context) {
	short_url_param := c.Param("short_url")

	join_with_base := os.Getenv("BASE_URL") + "/" + short_url_param

	link, status := link_service.GetInstance().GetLinkByShortUrl(join_with_base)

	if status.Code != http.StatusOK {
		c.JSON(int(status.Code), json_response.New(status.Code, status.Message, nil))
	}

	c.Redirect(int(status.Code), link.OriginalUrl)
}

func GetLinksByUserId(c *gin.Context) {
	user_id_body, _ := c.Get("user_id")

	user_id := user_id_body.(string)

	deleted_link, status := link_service.GetInstance().GetLinksByUserId(user_id)

	if status.Code != http.StatusCreated {
		c.JSON(int(status.Code), json_response.New(status.Code, status.Message, nil))
	}

	c.JSON(int(status.Code), json_response.New(status.Code, status.Message, deleted_link))
}
