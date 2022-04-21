package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sdetAutomation/golang-testing/models"
)

func ApiGetRequest(url string) []byte {

	fmt.Println("requested url: " + url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func GetGitHubLogin(data []byte) string {

	githubPayload := models.GitHubUserInfo{}

	err := json.Unmarshal(data, &githubPayload)
	if err != nil {
		fmt.Println(err)
	}

	return githubPayload.Login
}

func main() {
	apiResponse := ApiGetRequest("https://api.github.com/users/sdetAutomation")
	fmt.Println(" ")
	fmt.Println("*****************")
	fmt.Println(string(apiResponse))
	fmt.Println("*****************")
	fmt.Println(" ")

	login := GetGitHubLogin(apiResponse)
	
	fmt.Println("login : " + login)
}
