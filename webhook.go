package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Webhook struct {
	Username  string `json:"username,omitempty"`
	Content   string `json:"content,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Flags     int    `json:"flags,omitempty"`
}

func CreateWebhook() Webhook {
	return Webhook{
		Username:  "",
		Content:   "",
		AvatarUrl: "",
		Flags:     0,
	}
}

func (w Webhook) SetContent(content string) Webhook {
	w.Content = content
	return w
}

func (w Webhook) SetUsername(username string) Webhook {
	w.Username = username
	return w
}

func (w Webhook) SetAvatarUrl(avatarUrl string) Webhook {
	w.AvatarUrl = avatarUrl
	return w
}

func (w Webhook) SuppressEmbeds() Webhook {
	w.Flags = w.Flags | 4
	return w
}

func (w Webhook) Post(webhookUrl string) (Webhook, error) {
	data, err := json.Marshal(w)
	if err != nil {
		return w, fmt.Errorf("failed to post webhook: %w", err)
	}
	resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return w, fmt.Errorf("failed to post webhook: %w", err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	defer resp.Body.Close()
	return w, nil
}

func PostBuildingIdea(idea BuildingIdea, config Config) {
	// Create a webhook
	var link string
	if len(idea.SourceUrl) > 0 {
		link = fmt.Sprintf("[%s](%s)", idea.SourceTitle, idea.SourceUrl)
	} else {
		link = idea.SourceTitle
	}
	_, err := CreateWebhook().SuppressEmbeds().
		SetContent(fmt.Sprintf(`
<@&1407712465868820530>
This Week's Building Idea
# %s
-# Source: %s
`,
			idea.Title,
			link,
		)).
		SetUsername(
			"Minecraft Building Ideas",
		).
		Post(config.WebhookUrl)
	if err != nil {
		log.Fatal(err)
	}
}
