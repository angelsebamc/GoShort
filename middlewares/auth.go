package middlewares

import (
	"goshort/repositories/user_repository"
	"goshort/utils/json_response"
	"goshort/utils/jwt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")

		if email == nil {
			c.AbortWithStatusJSON(403, json_response.New(403, "unauthorized", nil))
			return
		}

		request_jwt := c.Request.Header.Get("Authorization")
		if request_jwt == "" {
			c.AbortWithStatusJSON(403, json_response.New(403, "invalid token", nil))
			return
		}

		jwt_claims, claims_err := jwt.GetInstance().ExtractTokenClaims(request_jwt)

		if claims_err != nil {
			c.AbortWithStatusJSON(500, json_response.New(500, claims_err.Error(), nil))
			return
		}

		email_claim := jwt_claims["email"].(string)
		user_id_claim := jwt_claims["id"].(string)

		if email != email_claim {
			c.AbortWithStatusJSON(403, json_response.New(403, "invalid user", nil))
			return
		}

		user_object_id, err := primitive.ObjectIDFromHex(user_id_claim)

		if err != nil {
			c.AbortWithStatusJSON(403, json_response.New(403, "invalid user", nil))
		}

		user_from_db := user_repository.GetInstance().GetUserById(user_object_id)

		if user_from_db == nil {
			c.AbortWithStatusJSON(403, json_response.New(404, "user does not exists", nil))
		}

		c.Set("user_email", email_claim)
		c.Set("user_id", user_from_db.ID)

		c.Next()
	}
}
