package shortener_logic

import (
	"encoding/base64"
	"math/rand"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const short_url_length = 6

func GenerateRandomString() string {
	random_string := make([]byte, short_url_length)
	for i := range random_string {
		random_string[i] = charset[rand.Intn(len(charset))]
	}

	random_string_b64 := base64.URLEncoding.EncodeToString(random_string)

	return string(random_string_b64)
}

func CreateShortURL(orignal_url string) string {
	base_url := os.Getenv("BASE_URL")
	short_url := GenerateRandomString()

	return base_url + "/" + short_url
}
