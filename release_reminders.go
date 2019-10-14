package main

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/logger"
)

func main() {
	logger.Instance().Printf("Projects repositories config: %s\n", common.ProjectRepConfig)
	logger.Instance().Printf("Projects repositories values: %s\n", common.ProjectsRepositoriesValues)
}
