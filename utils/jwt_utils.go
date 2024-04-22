package utils

import (
	"fmt"
	"goshort/models"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SecretKey []byte
}

var (
	jwt_instance *JWT
	once         sync.Once
)

func GetJwtUtilsInstance() *JWT {
	once.Do(func() {
		secretKey := []byte(os.Getenv("JWT_SECRET"))
		jwt_instance = &JWT{SecretKey: secretKey}
	})
	return jwt_instance
}

func (j *JWT) GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(j.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWT) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		return fmt.Errorf("error trying to parse jwt token")
	}

	if !token.Valid {
		return fmt.Errorf("the token is not valid")
	}

	return nil
}
