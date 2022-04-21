package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// httptest mock server response
func Test_tc0001_ApiGetRequest(t *testing.T) {

	var td_mocked_response = `{"my_mock_response":"my_mocked_body"}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(td_mocked_response))
	}))
	defer server.Close()

	// fmt.Println(server.URL)

	baseUrl := server.URL

	value := ApiGetRequest(baseUrl)

	assert.Equal(t, td_mocked_response, string(value))
}

// unit test business logic / GitHubLogin function
func Test_tc0002_GitHubLogin(t *testing.T) {

	td_expected_response := "sdetAutomation"

	td_payload := "{\"login\": \"sdetAutomation\"}"
	rawIn := json.RawMessage(td_payload)

	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		fmt.Println(err)
	}

	returnValue := GetGitHubLogin(bytes)

	// fmt.Println(returnValue)

	assert.Equal(t, td_expected_response, returnValue)
}
