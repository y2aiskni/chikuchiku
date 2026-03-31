package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WebHookBody struct {
	Content *string  `json:"content,omitempty"`
	Embeds  []*Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	URL         string      `json:"url,omitempty"`
	Footer      EmbedFooter `json:"footer,omitempty"`
}

type EmbedFooter struct {
	Text string `json:"text"`
}

func PostWebhook(url string, v WebHookBody) error {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to json.Marshal(): %w", err)
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to http.Post(): %w", err)
	}

	return nil
}
