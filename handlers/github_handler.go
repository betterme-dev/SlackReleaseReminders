package handlers

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/logger"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"sort"
	"strings"
	"unicode"
)

type (
	GitHubConfig struct {
		Token string
	}
	GitHubRepoReleases struct {
		RepositoryName string
		ReleasesTags   []string
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

func FetchRepositoriesReleases(repositoriesNames *[]string) *[]GitHubRepoReleases {
	repositoriesReleases := make([]GitHubRepoReleases, 0)
	// Loop through all required repositories
	for _, repoName := range *repositoriesNames {
		// Request list of releases for each of repository
		repoReleases, _, err := gitHubClient.Repositories.ListReleases(context.Background(), organizationName, repoName, &github.ListOptions{})
		if err != nil {
			logger.Instance().Errorf("Failed to fetch releases for repository: %s with error %s\n", repoName, err)
		}
		// Sort list by published time (just in case they ain't sorted)
		sort.Slice(repoReleases, func(i, j int) bool {
			return repoReleases[i].PublishedAt.Time.Before(repoReleases[j].PublishedAt.Time)
		})

		// Iterate over all repository releases and grab releases tags names, extracting versions from them
		tagsNames := make([]string, 0)
		for _, release := range repoReleases {
			tagsNames = append(tagsNames, strings.TrimFunc(*release.TagName, func(r rune) bool {
				return !unicode.IsNumber(r)
			}))
		}

		// If there are less than VersionToCheck in the project - take all of them
		if len(tagsNames) <= common.VersionToCheck {
			// Collect results into Repository name + list of git tags names structure
			repositoriesReleases = append(repositoriesReleases, GitHubRepoReleases{
				RepositoryName: repoName,
				ReleasesTags:   tagsNames,
			})
		} else {
			// Collect results into Repository name + list of git tags names structure
			repositoriesReleases = append(repositoriesReleases, GitHubRepoReleases{
				RepositoryName: repoName,
				ReleasesTags:   tagsNames[len(tagsNames)-common.VersionToCheck:],
			})
		}
	}

	return &repositoriesReleases
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
