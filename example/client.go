package example

import (
	"github.com/assetto-io/request/httpgo"
	"time"
)

var (
	httpClient httpgo.HttpClient
	mockKeeper httpgo.MockKeeper
)

func setHttpClient() {
	httpClient = httpgo.NewBuilder().
		SetResponseTimeout(3 * time.Second).
		SetMaxIdleConnections(5).
		Build()
}

func setMockHttpClient() {
	httpClient, mockKeeper = httpgo.NewBuilder().
		SetResponseTimeout(3 * time.Second).
		SetMaxIdleConnections(5).
		BuildMockClient()

	setMocks()
}
