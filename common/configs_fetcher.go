package common

import (
	"SlackReleaseReminders/logger"
	"github.com/spf13/viper"
)

type (
	ProjectRepositoryConfig struct {
		ProjectKey     string `mapstructure:"project"`
		RepositoryName string `mapstructure:"repository"`
	}
	ProjectsRepositories struct {
		ProjectsKeys      []string
		RepositoriesNames []string
	}
)

var ProjectRepConfig *[]ProjectRepositoryConfig
var ProjectsRepositoriesValues *ProjectsRepositories

func init() {
	ProjectRepConfig = readConfigs()
	ProjectsRepositoriesValues = mapToProjectsRepositoriesSlices(ProjectRepConfig)
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

func readConfigs() *[]ProjectRepositoryConfig {
	// setup viper
	viper.Reset()
	viper.SetConfigType("yaml")
	//viper.SetConfigName(os.Getenv("PROJECTS_REPOSITORIES_CONFIG"))
	viper.SetConfigName("android_projects_repositories")
	viper.AddConfigPath("configs")
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

	return conf
}
