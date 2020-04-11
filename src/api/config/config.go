package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
	maxNumberOfRoutinePerRequest = 10
)

func GetGithubAccessToke()  string {
	return githubAccessToken
}

func GetMaxRoutinesPerRequest() int {
	return maxNumberOfRoutinePerRequest
}