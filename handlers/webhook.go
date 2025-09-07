package handlers

import (
	"HookDeploy/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

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
	Repository   string    `yaml:"repository"`
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

	// Parse JSON body to check repository information
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		utils.WriteLog("Error: cannot parse webhook payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Check if repository object exists and has a full_name property
	repo, repoExists := payload["repository"].(map[string]interface{})
	if !repoExists {
		utils.WriteLog("Error: repository information missing in payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Repository information missing"})
		return
	}

	fullName, fullNameExists := repo["full_name"].(string)
	if !fullNameExists {
		utils.WriteLog("Error: repository full_name missing in payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Repository full_name missing"})
		return
	}

	cfgPath := "configs/" + strings.ReplaceAll(fullName, "/", "_") + ".yaml"
	data, err := os.ReadFile(cfgPath)

	if err != nil {
		utils.WriteLog("Error: can't receive information from " + fullName)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("can't receive information from %s", fullName)})
		return
	}

	// Now you can use fullName variable which contains the repository full_name
	utils.WriteLog(fmt.Sprintf("Received webhook from repository: %s", fullName))

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

	go func() {
		var cfg Config
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatal(err)
		}

		utils.WriteLog(fmt.Sprintf("Deploy %s", cfg.Repository))
		logName := "logs/" + strings.ReplaceAll(cfg.Repository, "/", "_") + ".log"
		for _, c := range cfg.PushCommands {
			fmt.Printf("➡️  Running: %s\n", c.Name)

			cmd := exec.Command(c.Cmd, c.Args...)
			cmd.Dir = cfg.Directory
			output, err := cmd.CombinedOutput()
			utils.WriteLog(string(output), logName)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				utils.WriteLog(fmt.Sprintf("❌ Error: %v", err), logName)
				utils.WriteLog(fmt.Sprintf("❌ Failed execute: %s", cfg.Repository))
				break
			}
		}
	}()
}
