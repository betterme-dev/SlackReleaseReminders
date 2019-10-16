package differ

import (
	"SlackReleaseReminders/handlers"
	"SlackReleaseReminders/merger"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateDiffNotFound(t *testing.T) {
	prjVersions := createProjectRepositoryJiraVersions(5)
	repoReleases := createGitHubRepoReleases(5)

	assert.Empty(t, *CalculateDiff(prjVersions, repoReleases))
}

func TestCalculateDiffFound(t *testing.T) {
	prjVersions := createProjectRepositoryJiraVersions(5)
	repoReleases := createGitHubRepoReleases(4)

	assert.NotEmpty(t, *CalculateDiff(prjVersions, repoReleases))
}

func createProjectRepositoryJiraVersions(releasesCount int) *[]merger.ProjectRepositoryJiraVersions {
	data := make([]merger.ProjectRepositoryJiraVersions, 0, 6)

	for i := 1; i <= 6; i++ {
		latestVersions := make([]string, 0, releasesCount)
		for j := 1; j <= releasesCount; j++ {
			latestVersions = append(latestVersions, fmt.Sprintf("%d.%d.%d", j, j, j))
		}
		data = append(data, merger.ProjectRepositoryJiraVersions{
			ProjectKey:     fmt.Sprintf("Project key: %d", i),
			RepositoryName: fmt.Sprintf("Repository name: %d", i),
			JiraVersions:   latestVersions,
		})
	}

	return &data
}

func createGitHubRepoReleases(tagsCount int) *[]handlers.GitHubRepoReleases {
	data := make([]handlers.GitHubRepoReleases, 0, 6)

	for i := 1; i <= 6; i++ {
		latestVersions := make([]string, 0, tagsCount)
		for j := 1; j <= tagsCount; j++ {
			latestVersions = append(latestVersions, fmt.Sprintf("%d.%d.%d", j, j, j))
		}
		data = append(data, handlers.GitHubRepoReleases{
			RepositoryName: fmt.Sprintf("Repository name: %d", i),
			TaggedVersions: latestVersions,
		})
	}

	return &data
}
