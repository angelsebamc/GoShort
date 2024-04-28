package link_handler

import (
	"goshort/dtos/link_dto"
	"goshort/services/link_service"
	"goshort/utils/json_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewShortUrl(c *gin.Context) {
	link_body, exists := c.Get("link_body")

	if !exists {
		c.JSON(http.StatusInternalServerError, json_response.New(http.StatusNotFound, "url not given", nil))
	}

	original_link := link_body.(link_dto.LinkDTO_Post)

	new_link, status := link_service.GetInstance().CreateShortURL(c, &original_link)

	if status.Code != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, json_response.New(status.Code, status.Message, nil))
	}

	c.JSON(http.StatusCreated, json_response.New(status.Code, status.Message, new_link))
}
