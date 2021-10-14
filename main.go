package main

import (
	"fmt"
	"github/vonuki/github_api/api"
	"time"
)

func main() {
	fmt.Println("Github test api")

	apiInstance := api.GetCurrentApi()
	// fmt.Println("My API token = ", apiInstance.Token)

	response, err := apiInstance.GetUserRepos("vonuki")
	if err != nil {
		fmt.Println("Error happened = ", err)
	}

	// test of channels
	my_text := make(chan string)
	go print_me(my_text)
	my_text <- fmt.Sprint("Status = ", response.Status)
	my_text <- fmt.Sprint("Repo counts = ", len(response.Repos))
	for i := 0; i < len(response.Repos); i++ {
		time.Sleep(time.Second)
		my_text <- fmt.Sprint(response.Repos[i].Name)
	}
}

// test of channels
func print_me(text chan string) {
	for {
		fmt.Println(<-text)
	}
}
