package main

import (
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