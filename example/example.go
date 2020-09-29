package main

import (
	"fmt"
	"github.com/assetto-io/request/request"
	"net/http"
	"time"
)

var (
	httpClient = makeGithubClient()
)

func makeGithubClient() request.HttpClient {
	customHeaders := make(http.Header)
	customHeaders.Set("Authorization", "Bearer: ABC-123")

	client := request.NewBuilder().
		SetCommonHeaders(customHeaders).
		SetConnectionTimeout(1 * time.Second).
		SetResponseTimeout(3 * time.Second).Build()

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
