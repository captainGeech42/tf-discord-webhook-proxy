package main

// Config items in config.json
type Config struct {
	WebhookURL   string
	Port         int
	RichMessages bool
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

// DiscordWebhook is a partial struct for https://discord.com/developers/docs/resources/webhook#execute-webhook
type DiscordWebhook struct {
	Content string         `json:"content"`
	Embeds  []DiscordEmbed `json:"embeds"`
}

// DiscordEmbed is a partial struct for https://discord.com/developers/docs/resources/channel#embed-object
type DiscordEmbed struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	URL         string              `json:"url"`
	Timestamp   string              `json:"timestamp"` // needs to be ISO8601
	Color       int                 `json:"color"`     // 0xRRGGBB
	Fields      []DiscordEmbedField `json:"fields"`
}

// DiscordEmbedField is a struct for https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
