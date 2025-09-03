package handlers

import (
	"HookDeploy/utils"
	"fmt"
	"net/http"
	"os"

	"io"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	event := c.GetHeader("X-GitHub-Event")
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read body"})
		return
	}

	signature := c.GetHeader("X-Hub-Signature-256")
	verifSignatue := utils.GenerateSignature(os.Getenv("SHA256_SECRET"), body)

	if signature != verifSignatue {
		utils.WriteLog("Error: signature is not valid")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signature is not valid"})
		return
	}

	fmt.Printf("Event: %s\n", event)
	fmt.Printf("Payload: %s\n", string(body))

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
