package main

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/handlers"
	"SlackReleaseReminders/logger"
)

func main() {
	jPKeys := common.ProjectsRepositoriesValues.ProjectsKeys
	jiraVersions := handlers.RetrieveJiraVersionsByKeys(jPKeys)
	logger.Instance().Printf("Versions: %s\n", jiraVersions)
}
