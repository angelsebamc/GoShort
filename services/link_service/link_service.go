package link_service

import (
	"goshort/dtos/link_dto"
	shortener_logic "goshort/logic"
	"goshort/repositories/link_repository"
	"goshort/utils/http_status"

	"github.com/gin-gonic/gin"
)

type LinkService struct{}

var instance *LinkService

func GetInstance() *LinkService {
	if instance == nil {
		instance = &LinkService{}
	}
	return instance
}

// methods

// the user has to exist and be authenticated
// the short url has to be unique and the original url has to be valid
// if the original url has a short url, return the short url
func (ls *LinkService) CreateShortURL(c *gin.Context, link *link_dto.LinkDTO_Post) (*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {
	user_id_get, user_id_exists := c.Get("user_id")

	if !user_id_exists {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusBadRequest, Message: "user does not exist"}
	}

	user_id := user_id_get.(string)

	short_link := shortener_logic.CreateShortURL()

	new_link := &link_dto.LinkDTO_Info{
		ShortUrl:    short_link,
		OriginalUrl: link.OriginalUrl,
		UserID:      user_id,
	}

	new_link_db, err := link_repository.GetInstance().AddLink(new_link)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusBadRequest, Message: err.Error()}
	}

	return new_link_db, &http_status.HTTPStatus{Code: http_status.StatusCreated, Message: "short link created"}
}

func (ls *LinkService) GetLinkByOriginalUrl(original_url string) (*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {
	link := link_repository.GetInstance().GetLinkByOriginalUrl(original_url)

	if link == nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusNotFound, Message: "link not found"}
	}

	return link, &http_status.HTTPStatus{Code: http_status.StatusOK, Message: "link found"}
}

func (ls *LinkService) GetLinkByShortUrl(short_url string) (*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {
	link := link_repository.GetInstance().GetLinkByShortUrl(short_url)

	if link == nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusNotFound, Message: "link not found"}
	}

	return link, &http_status.HTTPStatus{Code: http_status.StatusOK, Message: "link found"}
}
