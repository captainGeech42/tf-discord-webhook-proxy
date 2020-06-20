package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tkanos/gonfig"
)

// global config
var config Config

func main() {
	// parse config

	// check if env var
	env := os.Getenv("TF_PROXY_ENV")
	if env == "YES" {
		// pull config from env vars
		config.Port = 8080
		config.WebhookURL = os.Getenv("TF_PROXY_WEBHOOK_URL")
		rich := os.Getenv("TF_PROXY_RICH_MESSAGES")
		if rich == "NO" {
			config.RichMessages = false
		} else {
			config.RichMessages = true
		}
	} else {
		// don't pull from env, look for config.json
		err := gonfig.GetConf("config.json", &config)
		if err != nil {
			log.Fatal("Failed to parse config: " + err.Error())
		}
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

	log.Println("Handling notifications for run " + webhook.RunID)
	sendDiscordMessage(webhook)
	log.Println("Finished handling notifications for run " + webhook.RunID)
}

func sendDiscordMessage(webhook TerraformWebhook) {
	red := 0xff0000
	green := 0x00ff00
	yellow := 0xedb021

	for _, n := range webhook.Notifications {
		var discordMsg DiscordWebhook

		if config.RichMessages {
			var embed DiscordEmbed
			embed.Title = "Terraform Status"
			embed.Description = fmt.Sprintf("A new Terraform notification has been sent:\n\n**%s**", n.Message)
			embed.URL = webhook.RunURL
			embed.URL = "https://terraform.io"

			if n.RunStatus == "planned_and_finished" {
				embed.Color = green
			} else if n.RunStatus == "errored" {
				embed.Color = red
			} else {
				// this includes "discarded" or any other field in
				// https://www.terraform.io/docs/cloud/api/run.html#run-states
				embed.Color = yellow
			}

			embed.Fields = []DiscordEmbedField{
				{
					Name:  "Run Status",
					Value: n.RunStatus,
				},
				{
					Name:  "Run Message",
					Value: webhook.RunMessage,
				},
				{
					Name:   "Run Updated By",
					Value:  n.RunUpdatedBy,
					Inline: true,
				},
				{
					Name:   "Run Created By",
					Value:  webhook.RunCreatedBy,
					Inline: true,
				},
			}

			discordMsg.Embeds = []DiscordEmbed{embed}
		} else {
			discordMsg.Content = n.Message
		}

		if !makeDiscordRequest(discordMsg) {
			log.Printf("Discord message failed to send for notification \"%s\"", n.Message)
		}
	}
}

func makeDiscordRequest(msg DiscordWebhook) bool {
	jsonBody, err := json.Marshal(msg)
	if err != nil {
		log.Println("Failed to marshal Discord webhook: " + err.Error())
		return false
	}

	resp, err := http.Post(config.WebhookURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Failed to make Discord webhook request: " + err.Error())
		return false
	}

	resp.Body.Close()
	return true
}
