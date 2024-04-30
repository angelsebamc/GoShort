package link_service

import (
	"goshort/dtos/link_dto"
	shortener_logic "goshort/logic"
	"goshort/repositories/link_repository"
	"goshort/utils/http_status"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (ls *LinkService) CreateLink(c *gin.Context, link *link_dto.LinkDTO_Post) (*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {
	//TODO: maybe i can get this from the handler
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

func (ls *LinkService) DeleteLinkById(link_id string) (*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {

	object_id, err := primitive.ObjectIDFromHex(link_id)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	link, err_link := link_repository.GetInstance().DeleteLinkById(object_id)

	if err_link != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err_link.Error()}
	}

	return link, &http_status.HTTPStatus{Code: http_status.StatusOK, Message: "link deleted"}
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

func (ls *LinkService) GetLinksByUserId(user_id string) ([]*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {

	object_id, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	links, err := link_repository.GetInstance().GetLinksByUserId(object_id)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	return links, &http_status.HTTPStatus{Code: http_status.StatusOK, Message: "retrieving links successfully"}
}
