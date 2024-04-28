package middlewares

import (
	"goshort/utils/json_response"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")

		if token == nil {
			c.AbortWithStatusJSON(403, json_response.New(403, "unauthorized", nil))
			return
		}

		c.Next()
	}
}
