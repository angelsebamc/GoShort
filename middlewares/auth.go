package middlewares

import (
	"goshort/repositories/user_repository"
	"goshort/utils/json_response"
	"goshort/utils/jwt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		email := session.Get("email")

		if email == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "unauthorized", nil))
			return
		}

		request_bearer_jwt := c.Request.Header.Get("Authorization")
		if request_bearer_jwt == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "invalid token", nil))
			return
		}

		request_jwt := strings.Split(request_bearer_jwt, " ")

		if len(request_jwt) != 2 || request_jwt[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "invalid token", nil))
			return
		}

		jwt_claims, claims_err := jwt.GetInstance().ExtractTokenClaims(request_jwt[1])

		if claims_err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, json_response.New(http.StatusInternalServerError, claims_err.Error(), nil))
			return
		}

		email_claim := jwt_claims["email"].(string)
		user_id_claim := jwt_claims["id"].(string)

		if email != email_claim {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "invalid user", nil))
			return
		}

		user_object_id, err := primitive.ObjectIDFromHex(user_id_claim)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusForbidden, "invalid user", nil))
			return
		}

		user_from_db, err := user_repository.GetInstance().GetUserById(user_object_id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusInternalServerError, "error trying to get the user", nil))
			return
		}

		if user_from_db == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, json_response.New(http.StatusNotFound, "user does not exists", nil))
			return
		}

		c.Set("user_email", email_claim)
		c.Set("user_id", user_from_db.ID)

		c.Next()
	}
}
