package user_handler

import (
	"goshort/dtos/user_dto"
	"goshort/services/user_service"
	"goshort/utils/json_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	var new_user user_dto.UserDTO_Registration

	if err := c.BindJSON(&new_user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	create_user, status := user_service.GetInstance().CreateUser(&new_user)

	if status.Code != http.StatusCreated {
		c.JSON(int(status.Code), json_response.New(status.Code, status.Message, nil))
		return
	}

	c.JSON(int(status.Code), json_response.New(status.Code, status.Message, create_user))
}

func Auth(c *gin.Context) {
	var user user_dto.UserDTO_Login

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, json_response.New(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	token, status := user_service.GetInstance().GetUserByEmail(user.Email)

	if status.Code != http.StatusOK {
		c.JSON(int(status.Code), json_response.New(status.Code, status.Message, nil))
		return
	}

	user_with_token := &user_dto.UserDTO_Info_Token{
		Username: token.Username,
		Email:    token.Email,
		Token:    token.Token,
	}

	c.JSON(int(status.Code), json_response.New(status.Code, status.Message, user_with_token))
}
