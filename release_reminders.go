package main

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/differ"
	"SlackReleaseReminders/handlers"
	"SlackReleaseReminders/logger"
	"SlackReleaseReminders/merger"
	"os"
)

func main() {
	// Fetch configs for jira project key, repository pairs
	configs, groupedConfigs := common.FetchConfigs()
	// Fetch latest jira versions
	jiraVersions := handlers.RetrieveJiraVersionsByKeys(groupedConfigs.ProjectsKeys)
	// Merge fetched configs with latest jira versions
	mergedResults := merger.MergeProjectRepositoryConfigWithJiraVersions(configs, jiraVersions)
	// Fetch latest repositories versions
	repositoriesReleases := handlers.FetchRepositoriesReleasesByRepoNames(&groupedConfigs.RepositoriesNames)
	// There is a probably mismatch between Jira projects count and scanned repos count
	if len(*mergedResults) != len(*repositoriesReleases) {
		logger.Instance().Errorf("Mismatch between Jira projects count and scanned repos count!")
		os.Exit(1)
	}
	// Calculate diff, send slack alarm if needed
	diffResult := differ.CalculateDiff(mergedResults, repositoriesReleases)
	// If diff found - loop through the all missed versions and send notification
	if len(*diffResult) > 0 {
		for _, repo := range *diffResult {
			for _, version := range repo.MissedVersions {
				handlers.SendSlackAlarm(repo.RepoName, version)
			}
		}
	}
}
