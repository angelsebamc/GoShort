package main

import (
	"fmt"
	"goshort/routes/link_routes.go"
	"goshort/routes/user_routes"
	"goshort/utils/json_response"
	"net/http"
	"os"

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

	//not found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, json_response.New(http.StatusNotFound, "not found", nil))
	})

	//setup routes
	user_routes.SetupRoutes(router)
	link_routes.SetupRoutes(router)

	router.Run(os.Getenv("PORT"))
}
