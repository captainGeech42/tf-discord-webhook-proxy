package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tkanos/gonfig"
)

// Config items in config.json
type Config struct {
	WebhookURL string
	Port       int
}

// TerraformWebhook is the main webhook payload from TF
type TerraformWebhook struct {
	Version          int                     `json:"payload_version"`
	ConfigID         string                  `json:"notification_configuration_id"`
	RunURL           string                  `json:"run_url"`
	RunID            string                  `json:"run_id"`
	RunMessage       string                  `json:"run_message"`
	RunCreatedAt     string                  `json:"run_created_at"`
	RunCreatedBy     string                  `json:"run_created_by"`
	WorkspaceID      string                  `json:"workspace_id"`
	WorkspaceName    string                  `json:"workspace_name"`
	OrganizationName string                  `json:"organization_name"`
	Notifications    []TerraformNotification `json:"notifications"`
}

// TerraformNotification is the notification in the TF webhook
type TerraformNotification struct {
	Message      string `json:"message"`
	Trigger      string `json:"trigger"`
	RunStatus    string `json:"run_status"`
	RunUpdatedAt string `json:"run_updated_at"`
	RunUpdatedBy string `json:"run_updated_by"`
}

// global config
var config Config

func main() {
	// parse config
	err := gonfig.GetConf("config.json", &config)
	if err != nil {
		log.Fatal("Failed to parse config: " + err.Error())
	}

	log.Printf("Using webhook URL: %s\n", config.WebhookURL)

	http.HandleFunc("/webhook", handleIncomingWebhook)

	log.Printf("Listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}

func handleIncomingWebhook(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to get incoming webhook body: " + err.Error())
		return
	}

	var webhook TerraformWebhook
	err = json.Unmarshal(buf, &webhook)
	if err != nil {
		log.Println("Failed to unmarshal Terraform webhook: " + err.Error())
		return
	}

	sendDiscordMessage(webhook)
}

func sendDiscordMessage(webhook TerraformWebhook) {
	for _, s := range webhook.Notifications {
		log.Println(s.Message)
	}
}
