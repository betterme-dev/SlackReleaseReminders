package merger

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/handlers"
)

type (
	ProjectRepositoryJiraVersions struct {
		ProjectKey     string
		RepositoryName string
		JiraVersions   []string
	}
)

// Merges Jira Project Key, GitHub Repository Name and Jira latest release versions
func MergeProjectRepositoryConfigWithJiraVersions(prConfigs *[]common.ProjectRepositoryConfig, jVersions []*handlers.JiraRecentVersions) *[]ProjectRepositoryJiraVersions {
	mergedResult := make([]ProjectRepositoryJiraVersions, 0)

	for _, prConfig := range *prConfigs {
		for _, jv := range jVersions {
			if prConfig.ProjectKey == jv.ProjectKey {
				mergedResult = append(mergedResult, ProjectRepositoryJiraVersions{
					ProjectKey:     prConfig.ProjectKey,
					RepositoryName: prConfig.RepositoryName,
					JiraVersions:   jv.LatestVersions,
				})
			}
		}
	}

	return &mergedResult
}
