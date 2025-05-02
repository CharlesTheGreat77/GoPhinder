package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordNotifier struct {
	WebhookURL string
}

// constructor for DiscordNotifier.
func NewDiscordNotifier(webhookURL string) *DiscordNotifier {
	return &DiscordNotifier{WebhookURL: webhookURL}
}

// send a notification to Discord.
func (d *DiscordNotifier) Notify(title string, fields map[string]string) error {
	var embedFields []map[string]interface{}
	for k, v := range fields {
		embedFields = append(embedFields, map[string]interface{}{
			"name":   k,
			"value":  v,
			"inline": false,
		})
	}

	payload := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       title,
				"color":       0x00ffcc, // teal color
				"fields":      embedFields,
				"description": "📡 GoPhinder Notification",
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("[-] failed to marshal embed payload: %v", err)
	}

	req, err := http.NewRequest("POST", d.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("[-] failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[-] failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("[-] failed to send notification, status code: %d", resp.StatusCode)
	}

	return nil
}
