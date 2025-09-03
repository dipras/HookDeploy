package handlers

import (
	"HookDeploy/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"io"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Command struct {
	Name string   `yaml:"name"`
	Cmd  string   `yaml:"cmd"`
	Args []string `yaml:"args,omitempty"`
}

type Config struct {
	Directory    string    `yaml:"directory"`
	PushCommands []Command `yaml:"on_push"`
}

func WebhookHandler(c *gin.Context) {
	event := c.GetHeader("X-GitHub-Event")
	body, err := io.ReadAll(c.Request.Body)

	if event != "push" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only push event avaible"})
		return
	}

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

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range cfg.PushCommands {
		fmt.Printf("➡️  Running: %s\n", c.Name)

		cmd := exec.Command(c.Cmd, c.Args...)
		cmd.Dir = cfg.Directory
		_, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			utils.WriteLog(fmt.Sprintf("❌ Error: %v", err))
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
