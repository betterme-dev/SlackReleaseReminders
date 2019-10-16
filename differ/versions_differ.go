package differ

import (
	"SlackReleaseReminders/handlers"
	"SlackReleaseReminders/merger"
)

type (
	RepoReleaseDiff struct {
		RepoName       string
		MissedVersions []string
	}
)

// Finds diffs between Jira projects releases and GitHub releases, fires send Slack call back if diff has been found
func CalculateDiff(prjVersions *[]merger.ProjectRepositoryJiraVersions, repositoriesReleases *[]handlers.GitHubRepoReleases) []RepoReleaseDiff {
	diffResult := make([]RepoReleaseDiff, 0)
	// Iterate over all jira project-repository versions configs
	for _, prjVersion := range *prjVersions {
	diffForRepoFound:
		// Iterate over all github releases
		for _, repositoryRelease := range *repositoriesReleases {
			// If names repos names not matched - skip iteration
			if prjVersion.RepositoryName != repositoryRelease.RepositoryName {
				continue
			}
			// Find missed git tags
			missedTags := findDifference(prjVersion.JiraVersions, repositoryRelease.TaggedVersions)
			// If slice is empty - diff not found, skip iteration
			if len(missedTags) == 0 {
				continue
			}
			diffResult = append(diffResult, RepoReleaseDiff{
				RepoName:       prjVersion.RepositoryName,
				MissedVersions: missedTags,
			})
			// Versions diff for this repo has been computed - should start with the next one
			break diffForRepoFound
		}
	}

	return diffResult
}

// Difference returns the elements in `a` that aren't in `b`, using maps under the hood,
// because maps are ~O(1) and here is an ~O(n)
func findDifference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}
