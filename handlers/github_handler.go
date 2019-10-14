package handlers

import (
	"SlackReleaseReminders/logger"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

type (
	GitHubConfig struct {
		Token string
	}
)

const (
	organizationName = "betterme-dev"
)

var gitHubConfigs *GitHubConfig
var gitHubClient *github.Client

func init() {
	gitHubConfigs = retrieveGitHubConfigs()
	gitHubClient = createGitHubClient()
}

func FetchRepositoriesByNames(repositoriesNames *[]string) *[]github.Repository {
	foundRepos := make([]github.Repository, 0)
	// Request all repos list in organisation
	repos, _, err := gitHubClient.Repositories.ListByOrg(context.Background(), organizationName, &github.RepositoryListByOrgOptions{})
	if err != nil {
		logger.Instance().Errorf("Failed to fetch repositories for organization: %s", err)
	}
	// Selecting those we are interested in
	for _, repoName := range *repositoriesNames {
		for _, repo := range repos {
			if repo.Name == &repoName {
				foundRepos = append(foundRepos, *repo)
			}
		}
	}

	return &foundRepos
}

// Make auth and instantiate
func createGitHubClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gitHubConfigs.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

// Retrieves jiraConfigs from environment, probably Jenkins
func retrieveGitHubConfigs() *GitHubConfig {
	return &GitHubConfig{
		Token: os.Getenv("GITHUB_TOKEN"),
	}
}
