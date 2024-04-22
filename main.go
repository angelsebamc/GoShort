package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	//Loading environment variables
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("Error loading .env")
	}

}
