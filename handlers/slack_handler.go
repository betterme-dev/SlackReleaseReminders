package handlers

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/logger"
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"os"
)

type (
	SlackConfig struct {
		WebHookUrl  string
		ChannelName string
	}
)

const gitHubUrlScheme = "https://github.com/" + common.OrganizationName + "/"

func SendSlackAlarm(repositoryName string, releaseVersion string) {
	// Get Slack config
	configs := retrieveSlackConfigs()

	// Format message - repo name + repo version
	message := fmt.Sprintf("Repository: %s not tagged with release verison: %s", repositoryName, releaseVersion)

	// Add repo url through the attachment
	attachment := slack.Attachment{}
	attachment.AddAction(slack.Action{Type: "button", Text: "RepoURL", Url: gitHubUrlScheme + repositoryName, Style: "danger"})

	// Construct payload
	payload := slack.Payload{
		Username:    "Release tag reminder",
		Text:        message,
		Channel:     configs.ChannelName,
		Attachments: []slack.Attachment{attachment},
	}

	// Try to send Slack message
	err := slack.Send(configs.WebHookUrl, "", payload)

	if err != nil {
		logger.Instance().Errorln("Failed to send slack alarm: %s", err)
	}
}

// Retrieves jiraConfigs from environment, probably Jenkins
func retrieveSlackConfigs() *SlackConfig {
	return &SlackConfig{
		WebHookUrl:  os.Getenv("SLACK_WEBHOOK_URL"),
		ChannelName: os.Getenv("SLACK_CHANNEL_NAME"),
	}
}
