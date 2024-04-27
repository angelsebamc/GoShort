package link_service

import (
	"goshort/dtos/link_dto"
	shortener_logic "goshort/logic"
	"goshort/repositories/link_repository"
	"goshort/repositories/user_repository"
	"goshort/utils/http_status"
	"goshort/utils/jwt"

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

// the user has to exist and be authenticated
// the short url has to be unique and the original url has to be valid
// if the original url has a short url, return the short url
func (ls *LinkService) CreateShortURL(c *gin.Context, link *link_dto.LinkDTO_Post) (*link_dto.LinkDTO_Get, *http_status.HTTPStatus) {
	request_jwt := c.Request.Header.Get("Authorization")
	if request_jwt == "" {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: "invalid token"}
	}

	jwt_claims, claims_err := jwt.GetInstance().ExtractTokenClaims(request_jwt)

	if claims_err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: claims_err.Error()}
	}

	//try to get user from jwt claims
	user_id := jwt_claims["id"].(string)

	user_object_id, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	user_from_db := user_repository.GetInstance().GetUserById(user_object_id)

	if user_from_db == nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusUnauthorized, Message: "user not found"}
	}

	//creates a new short link
	short_link := shortener_logic.CreateShortURL(link.OriginalUrl)

	new_link := &link_dto.LinkDTO_Info{
		ShortUrl:    short_link,
		OriginalUrl: link.OriginalUrl,
		UserID:      user_from_db.ID,
	}

	new_link_db, err := link_repository.GetInstance().AddLink(new_link)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusBadRequest, Message: err.Error()}
	}

	return new_link_db, &http_status.HTTPStatus{Code: http_status.StatusCreated, Message: "short link created"}
}
