package shortener_logic

import (
	"goshort/utils"
	"os"
)

const short_url_length = 6

func CreateShortURL(orignal_url string) string {
	base_url := os.Getenv("BASE_URL")
	short_url := utils.GenerateRandomString(short_url_length)

	return base_url + "/" + short_url
}
