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

const (
	gitHubUrlScheme = "https://github.com/" + common.OrganizationName + "/"
)

func SendSlackAlarm(repositoryName string, releaseVersion string) {
	// Get Slack config
	configs := retrieveSlackConfigs()

	// Construct payload
	payload := slack.Payload{
		Username:    "Release tag reminder",
		Text:        "",
		Channel:     configs.ChannelName,
		Attachments: []slack.Attachment{prepareAttachment(repositoryName, releaseVersion)},
	}

	// Try to send Slack message
	err := slack.Send(configs.WebHookUrl, "", payload)

	if err != nil {
		logger.Instance().Errorln("Failed to send slack alarm: %s", err)
	}
}

func prepareAttachment(repositoryName string, releaseVersion string) slack.Attachment {
	repoLink := gitHubUrlScheme + repositoryName
	messageColor := "danger"
	attachment := slack.Attachment{}
	attachment.Title = &repositoryName
	attachment.TitleLink = &repoLink
	attachment.Color = &messageColor
	attachmentFields := []*slack.Field{
		{
			Title: "Message: ",
			Value: fmt.Sprint("Forgot to Git Tag new release version 🤦🏻 ‍"),
			Short: false,
		},
		{
			Title: "Repository: ",
			Value: repositoryName,
			Short: false,
		},
		{
			Title: "Version in Jira: ",
			Value: releaseVersion,
			Short: false,
		},
	}
	attachment.Fields = attachmentFields

	return attachment
}

// Retrieves jiraConfigs from environment, probably Jenkins
func retrieveSlackConfigs() *SlackConfig {
	return &SlackConfig{
		WebHookUrl:  os.Getenv("SLACK_WEBHOOK_URL"),
		ChannelName: os.Getenv("SLACK_CHANNEL_NAME"),
	}
}
