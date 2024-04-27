package main

import (
	"fmt"
	"goshort/handlers/session_handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	//Loading environment variables
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("Error loading .env")
	}

	//user hanlders
	r.POST("/api/user/register", session_handler.Register)
	r.POST("/api/user/login", session_handler.Login)

	r.Run(":8080")
}
