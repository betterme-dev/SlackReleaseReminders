package main

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/handlers"
	"SlackReleaseReminders/merger"
)

func main() {
	configs, groupedConfigs := common.FetchConfigs()
	jiraVersions := handlers.RetrieveJiraVersionsByKeys(groupedConfigs.ProjectsKeys)
	repositoriesReleases := handlers.FetchRepositoriesReleases(&groupedConfigs.RepositoriesNames)
	mergedResults := merger.MergeProjectRepositoryConfigWithJiraVersions(configs, jiraVersions)

	for _, mr := range *mergedResults {
		for _, rr := range *repositoriesReleases {
			if mr.RepositoryName == rr.RepositoryName {

			}
		}
	}
}
