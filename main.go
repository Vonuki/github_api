package main

import (
	"fmt"
	"github/vonuki/github_api/api"
)

func main() {
	fmt.Println("Github test api")

	apiInstance := api.GetCurrentApi()
	// fmt.Println("My API token = ", apiInstance.Token)

	response, err := apiInstance.GetUserRepos("vonuki")
	if err != nil {
		fmt.Println("Error happened = ", err)
	}
	fmt.Println("Status = ", response.Status)
	fmt.Println("Repo counts = ", len(response.Repos))
	for i := 0; i < len(response.Repos); i++ {
		fmt.Println(response.Repos[i].Name)
	}
}
