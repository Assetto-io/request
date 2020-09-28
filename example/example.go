package main

import (
	"fmt"
	"github.com/assetto-io/request/request"
	"net/http"
)

var (
	httpClient = makeGithubClient()
)

func makeGithubClient() request.HttpClient {
	client := request.Client()

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer: ABC-123")
	client.SetCommonHeaders(headers)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	getUrls()
}

func getUrls() {
	//headers := make(http.Header)
	//headers.Set("Authorization", "Bearer: ABC-123")
	response, err := httpClient.Get("https://api.github.com", nil)

	if err != nil {

	}

	fmt.Println(response.StatusCode)

	//bytes, _ := ioutil.ReadAll(response.Body)
	//fmt.Printf(string(bytes))

}
