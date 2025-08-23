package main

import (
	"fmt"
	"log"

	discordWebhook "github.com/dozerokz/discord-webhook-go"
)

func PostBuildingIdea(idea BuildingIdea, config Config) {
	// Create a webhook
	var link string
	if len(idea.SourceUrl) > 0 {
		link = fmt.Sprintf("[%s](%s)", idea.SourceTitle, idea.SourceUrl)
	} else {
		link = idea.SourceTitle
	}
	webhook, err := discordWebhook.CreateWebhook(fmt.Sprintf(`
<@&1407712465868820530>
This Week's Building Idea
# %s
-# Source: %s
`,
		idea.Title,
		link,
	),
		"Minecraft Building Ideas",
		"",
	) // replace with actual image url (string)
	if err != nil {
		log.Fatal(err)
	}

	// Send the webhook
	err = discordWebhook.SendWebhook(config.WebhookUrl, webhook)
	if err != nil {
		log.Fatal(err)
	}
}
