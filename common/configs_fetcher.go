package common

import (
	"SlackReleaseReminders/logger"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

type (
	// Simply just a named map with Jira Project Key as key and GitHub Repository Name as value
	ProjectRepositoryConfig struct {
		ProjectKey     string `mapstructure:"project"`
		RepositoryName string `mapstructure:"repository"`
	}
	// Holder for grouped project keys and repositories names (used for fetching info from Jira and GitHub APIs)
	ProjectsRepositories struct {
		ProjectsKeys      []string
		RepositoriesNames []string
	}
)

const (
	VersionToCheck   = 4
	OrganizationName = "betterme-dev"
)

// Fetches configs Jira Project Key - Repository name pairs, separate values as slice
func FetchConfigs() (*[]ProjectRepositoryConfig, *ProjectsRepositories) {
	// setup viper
	viper.Reset()
	viper.SetConfigType("yaml")
	viper.AddConfigPath(getBasePath() + "/configs")
	//viper.SetConfigName(os.Getenv("PROJECTS_REPOSITORIES_CONFIG"))
	viper.SetConfigName("android_projects_repositories")
	// read config and check if any error occurs
	err := viper.ReadInConfig()
	if err != nil {
		logger.Instance().Errorf("Failed to read the configs: %s\n", err)
	}

	var conf *[]ProjectRepositoryConfig
	// unmarshal read configs to the struct
	err = viper.UnmarshalKey("projects-repositories", &conf)
	if err != nil {
		logger.Instance().Errorf("Unable to decode into config struct, %s\n", err)
	}

	return conf, mapToProjectsRepositoriesSlices(conf)
}

func mapToProjectsRepositoriesSlices(prc *[]ProjectRepositoryConfig) *ProjectsRepositories {
	projectsNames := make([]string, 0)
	repositoriesNames := make([]string, 0)

	// Iterate over configs and group items
	for _, c := range *prc {
		projectsNames = append(projectsNames, c.ProjectKey)
		repositoriesNames = append(repositoriesNames, c.RepositoryName)
	}

	pr := ProjectsRepositories{
		ProjectsKeys:      projectsNames,
		RepositoriesNames: repositoriesNames,
	}

	return &pr
}

func getBasePath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
