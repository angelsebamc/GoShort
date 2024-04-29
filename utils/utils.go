package utils

import (
	"encoding/base64"
	"math/rand"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateRandomString(length int) string {
	random_string := make([]byte, length)
	for i := range random_string {
		random_string[i] = charset[rand.Intn(len(charset))]
	}

	random_string_b64 := base64.URLEncoding.EncodeToString(random_string)

	return string(random_string_b64)
}

func ValidateURL(url string) bool {
	match, err := regexp.MatchString(`^https://www\..+`, url)

	if err != nil {
		return false
	}

	return match
}
