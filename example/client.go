package example

import (
	"github.com/assetto-io/request/request"
	"time"
)

var (
	httpClient request.HttpClient
	mockKeeper request.MockKeeper
)

func setHttpClient() {
	httpClient = request.NewBuilder().
		SetResponseTimeout(3 * time.Second).
		SetMaxIdleConnections(5).
		Build()
}

func setMockHttpClient() {
	httpClient, mockKeeper = request.NewBuilder().
		SetResponseTimeout(3 * time.Second).
		SetMaxIdleConnections(5).
		BuildMockClient()

	setMocks()
}
