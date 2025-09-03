package main

import (
	"HookDeploy/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":9000")
}
