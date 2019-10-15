package merger

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/handlers"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMergeProjectRepositoryConfigWithJiraVersions(t *testing.T) {
	projectsRepositoriesConfigs := createProjectsRepositoriesConfigs()
	jiraVersions := createJiraVersions()

	expectedResult := createMergedResults()

	result := MergeProjectRepositoryConfigWithJiraVersions(projectsRepositoriesConfigs, jiraVersions)

	assert.Equal(t, result, expectedResult)
}

func createProjectsRepositoriesConfigs() *[]common.ProjectRepositoryConfig {
	data := make([]common.ProjectRepositoryConfig, 0)

	for i := 0; i <= 5; i++ {
		data = append(data, common.ProjectRepositoryConfig{
			ProjectKey:     fmt.Sprintf("Project key: %d", i),
			RepositoryName: fmt.Sprintf("Repository name: %d", i),
		})
	}

	return &data
}

func createJiraVersions() []*handlers.JiraRecentVersions {
	data := make([]*handlers.JiraRecentVersions, 0)

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			latestVersions := make([]string, 0)
			latestVersions = append(latestVersions, fmt.Sprintf("%d.%d.%d", j, j, j))

			data = append(data, &handlers.JiraRecentVersions{
				ProjectKey:     fmt.Sprintf("Project key: %d", i),
				LatestVersions: latestVersions,
			})
		}
	}

	return data
}

func createMergedResults() *[]ProjectRepositoryJiraVersions {
	data := make([]ProjectRepositoryJiraVersions, 0)

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			latestVersions := make([]string, 0)
			latestVersions = append(latestVersions, fmt.Sprintf("%d.%d.%d", j, j, j))

			data = append(data, ProjectRepositoryJiraVersions{
				ProjectKey:     fmt.Sprintf("Project key: %d", i),
				RepositoryName: fmt.Sprintf("Repository name: %d", i),
				JiraVersions:   latestVersions,
			})
		}
	}

	return &data
}
