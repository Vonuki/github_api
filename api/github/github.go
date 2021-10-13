package github

import (
	"encoding/json"
	"fmt"
	"github/vonuki/github_api/api/response"
	"io/ioutil"
	"net/http"
)

type GitHubApi struct {
	url   string
	Token string
}

func NewGitHubApi(token string) *GitHubApi {
	api := GitHubApi{}
	api.Token = token
	api.url = "https://api.github.com/users/%s/repos"
	return &api
}

func (a *GitHubApi) GetUserRepos(userName string) (response.GitResponse, error) {

	response := response.GitResponse{}
	// response.Repos = make([]api.Repo, 0, 10)

	resp, err := http.Get(fmt.Sprintf(a.url, userName))
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response.Repos)
	if err != nil {
		return response, err
	}

	response.Status = "succces"
	return response, nil

}
