package main

import (
	"fmt"
	session_routes "goshort/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	//Loading environment variables
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println("Error loading .env")
	}

	session_routes.SetupRoutes(router)

	router.Run(":8080")
}
