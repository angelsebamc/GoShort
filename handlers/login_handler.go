package handlers

//TODO: Login Handler

import (
	"goshort/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func Register(c *gin.Context) {
	var newUser dtos.UserDTO_Registration

	if err := c.BindJSON(&newUser); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// newUserDb := models.User{
	// 	Username: newUser.Username,
	// 	Email: newUser.Email,
	// 	Password: newUser.Password,
	// }
}
