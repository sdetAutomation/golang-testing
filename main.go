package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
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

 func main() {
	 apiResponse := ApiGetRequest("https://api.github.com/users/sdetAutomation")
	 fmt.Println(string(apiResponse))
 }