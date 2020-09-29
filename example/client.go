package example

import (
	"github.com/assetto-io/request/request"
	"time"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() request.HttpClient {
	client := request.NewBuilder().
		SetResponseTimeout(3 * time.Second).
		SetMaxIdleConnections(5).
		Build()

	return client
}
