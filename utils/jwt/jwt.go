package jwt

import (
	"fmt"
	"goshort/dtos/user_dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secretKey []byte
}

var instance *JWT

func GetInstance() *JWT {

	if instance == nil {
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		instance = &JWT{secretKey: secretKey}
	}

	return instance
}

func (j *JWT) GenerateToken(user *user_dto.UserDTO_Info) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(j.secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return fmt.Errorf("error trying to parse jwt token")
	}

	if !token.Valid {
		return fmt.Errorf("the token is not valid")
	}

	return nil
}

func (j *JWT) ExtractTokenClaims(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error trying to parse jwt token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("the token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("error trying to parse jwt token claims")
	}

	return claims, nil
}
