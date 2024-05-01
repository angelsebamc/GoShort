package session_handler

//TODO: Login Handler

import (
	"goshort/dtos/user_dto"
	"goshort/services/user_service"
	"goshort/utils/json_response"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var new_user user_dto.UserDTO_Registration

	if err := c.BindJSON(&new_user); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	create_user, status := user_service.GetInstance().CreateUser(&new_user)

	if status.Code != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, json_response.New(status.Code, status.Message, nil))
		return
	}

	c.JSON(http.StatusCreated, json_response.New(status.Code, status.Message, create_user))
}

func Login(c *gin.Context) {
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

	session := sessions.Default(c)
	session.Set("email", token.Email)

	session.Save()

	user_with_token := &user_dto.UserDTO_Info_Token{
		Username: token.Username,
		Email:    token.Email,
		Token:    token.Token,
	}

	c.JSON(int(status.Code), json_response.New(status.Code, status.Message, user_with_token))
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("email")
	session.Save()

	c.JSON(http.StatusOK, json_response.New(200, "Logged out", nil))
}
