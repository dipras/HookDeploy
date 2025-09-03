package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	name := c.Query("name") // ?name=John
	if name == "" {
		name = "World"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}
