package main

import (
	"fmt"
	"goshort/routes/link_routes.go"
	"goshort/routes/session_routes"
	"goshort/utils"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	//setup session
	store := cookie.NewStore([]byte(utils.GenerateRandomString(64)))
	router.Use(sessions.Sessions("goshort", store))

	//setup routes
	session_routes.SetupRoutes(router)
	link_routes.SetupRoutes(router)

	router.Run(os.Getenv("PORT"))
}
