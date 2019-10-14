package handlers

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type (
	GitHubConfig struct {
		Token         string
		ObservedRepos []string
	}
)

var gitHubConfigs *GitHubConfig
var gitHubClient *github.Client

func init() {
	gitHubConfigs = retrieveGitHubConfigs()
	gitHubClient = createGitHubClient()
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
		Token:         os.Getenv("GITHUB_TOKEN"),
		ObservedRepos: strings.Fields(os.Getenv("OBSERVED_REPOS")),
	}
}
