# The purpose:

Tool that compares Jira Release Versions and GitHub releases and if there is a mismatch - sends notification to the 
specified Slack channel.

In other words - it reminds to Git tag release version of the source code.

Release Version in Jira should be with status **Released**. Tool will compare last **4** Jira/GitHub versions.


## How to setup:

The idea that tool will be triggered by Jenkins through the CRON job.

There are a bunch of the environment variables that can be used for specific configuration.


#### Slack configuration:

```SLACK_WEBHOOK_URL``` - webhook url, that you can configure through the team's Settings/Apps/Incoming Webhook section.
```SLACK_CHANNEL_NAME``` - channel name in which you want to receive notifications.

#### Jira configuration:

```JIRA_TOKEN``` - token to make an auth.

```JIRA_USERNAME``` - username (could be an email) to make an auth.

```JIRA_PROJECT_URL``` - should be https://newsiteam.atlassian.net by default.

#### GitHub configuration:

```GITHUB_TOKEN``` - A GitHub token with **admin:org** access (is needed to access organization's repositories).

#### Jira project - GitHub repository mapping config:

We need provide a _.yml_ mapping file to understand which Jira Project Key (WLA, WIOS, MEDA and so on) corresponds to 
which GitHub repository name.

This file should be placed under **common/configs** folder with the following structure:

```
projects-repositories:
   - project: 'WLA'
     repository: 'ANDROID_BETTERME'
   - project: 'MEDA'
     repository: 'Meditation-Android'
   - project: 'MENA'
     repository: 'BetterMen-Android'
   - project: 'WA'
     repository: 'Walking-Android'
   - project: 'RA'
     repository: 'Running-Android'
```

```PROJECTS_REPOSITORIES_CONFIG``` - the environment variable that corresponds to the config file name, that should be
provided through the Jenkins.

#### Docker deployment:

Docker image deployed to the Docker hub: ```betterme/slack-release-reminders:1.0.1```.

Ask DevOps team to access the image if needed.

#### Jenkins:

You can use a pipeline job which is scheduled to run periodically using CRON. An example of running app inside the 
Docker container:

```
node('node-name') {
    stage('Check release versions') {
        docker.image('betterme/slack-release-reminders:1.0.7').inside("-e SLACK_CHANNEL_NAME=${SLACK_CHANNEL_NAME} \
        -e SLACK_WEBHOOK_URL=${SLACK_WEBHOOK_URL} \
        -e JIRA_USERNAME=${JIRA_USERNAME} \
        -e JIRA_TOKEN=${JIRA_TOKEN} \
        -e JIRA_PROJECT_URL=${JIRA_PROJECT_URL} \
        -e PROJECTS_REPOSITORIES_CONFIG=${PROJECTS_REPOSITORIES_CONFIG} \
        -e GITHUB_TOKEN=${GITHUB_TOKEN}") {
             sh "/go/bin/app"
        }
    }
}
```
All parameters can be injected through the *General* -> *This project is parametrized* settings.