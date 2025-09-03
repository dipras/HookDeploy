package main

import (
	"HookDeploy/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
