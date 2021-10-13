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

	responseInst := response.GitResponse{}
	// responseInst.Repos = make([]response.Repo, 0, 10)

	resp, err := http.Get(fmt.Sprintf(a.url, userName))
	if err != nil {
		return responseInst, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseInst, err
	}

	err = json.Unmarshal(body, &responseInst.Repos)
	if err != nil {
		return responseInst, err
	}

	responseInst.Status = "succces"
	return responseInst, nil

}
