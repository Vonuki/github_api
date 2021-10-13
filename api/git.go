package api

import (
	"github/vonuki/github_api/api/github"
	"github/vonuki/github_api/api/response"
)

type RepoApi interface {
	GetUserRepos(userName string) (response.GitResponse, error)
}

func GetCurrentApi() RepoApi {
	apiInstance := github.NewGitHubApi("some token")
	return apiInstance
}
