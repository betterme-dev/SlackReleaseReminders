package handlers

import (
	"SlackReleaseReminders/common"
	"SlackReleaseReminders/logger"
	"github.com/andygrunwald/go-jira"
	"os"
	"sort"
	"strings"
	"unicode"
)

type (
	JiraConfig struct {
		Token    string // Token with projects-level access
		Username string // Username (email)
		TeamUrl  string // Url of the team (base url)
	}
	JiraRecentVersions struct {
		ProjectKey     string   // WLA, MEDA, WIOS and so on
		LatestVersions []string // Size is equal to lastNVersions
	}
)

var jiraClient *jira.Client

func init() {
	jiraClient = createJiraClient()
}

// Fetches projects releases by passed Jira projects keys
func RetrieveJiraVersionsByKeys(jiraKeys []string) []*JiraRecentVersions {
	lastVersions := make([]*JiraRecentVersions, 0)
	projects := findProjectsByKeys(jiraKeys)
	for _, p := range projects {
		lastVersions = append(lastVersions, extractVersionsForProject(p))
	}
	return lastVersions
}

// Retrieves releases versions from specified project
func extractVersionsForProject(project *jira.Project) *JiraRecentVersions {
	// First we need to be sure that versions sorted by release date
	allVersions := project.Versions
	sort.Slice(allVersions, func(i, j int) bool {
		return allVersions[i].ReleaseDate <= allVersions[j].ReleaseDate
	})
	// Take into account only versions that are released, but not archived
	releasedVersions := make([]jira.Version, 0)
	for _, version := range project.Versions {
		if version.Released && !version.Archived {
			releasedVersions = append(releasedVersions, version)
		}
	}

	// Extract version names
	versionNames := make([]string, 0)
	for _, version := range releasedVersions {
		versionNames = append(versionNames, strings.TrimFunc(version.Name, func(r rune) bool {
			return !unicode.IsNumber(r)
		}))
	}

	// If there are less than VersionToCheck in the project - take all of them
	if len(versionNames) <= common.VersionToCheck {
		return &JiraRecentVersions{
			ProjectKey:     project.Key,
			LatestVersions: versionNames,
		}
	} else {
		// Otherwise take last 4 versions
		return &JiraRecentVersions{
			ProjectKey:     project.Key,
			LatestVersions: versionNames[len(versionNames)-common.VersionToCheck:],
		}
	}
}

// Retrieves projects by specified keys
func findProjectsByKeys(keys []string) []*jira.Project {
	teamProjects := make([]*jira.Project, 0)

	for _, key := range keys {
		// Try to retrieve project by the provided key
		project, _, err := jiraClient.Project.Get(key)
		if err != nil {
			logger.Instance().Errorf("Failed to retrieve project details by key: %s with error: %s\n", key, err)
		}
		teamProjects = append(teamProjects, project)
	}

	return teamProjects
}

// Make auth and instantiate
func createJiraClient() *jira.Client {
	// Retrieve configs first
	jc := retrieveJiraConfigs()
	// Init auth transport
	tp := jira.BasicAuthTransport{
		Username: jc.Username,
		Password: jc.Token,
	}
	// Try to init client
	client, err := jira.NewClient(tp.Client(), jc.TeamUrl)
	if err != nil {
		logger.Instance().Errorf("Failed to init Jira client: %s\n", err)
	}
	return client
}

// Retrieves jiraConfigs from environment, probably Jenkins
func retrieveJiraConfigs() *JiraConfig {
	return &JiraConfig{
		Token:    os.Getenv("JIRA_TOKEN"),
		Username: os.Getenv("JIRA_USERNAME"),
		TeamUrl:  os.Getenv("JIRA_PROJECT_URL"),
	}
}
