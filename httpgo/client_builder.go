package httpgo

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	client             *http.Client
	userAgent          string
	timeoutsDisabled   bool
}

// ClientBuilder provides a clean way to configure HTTP Client based on
// the combination of methods
type ClientBuilder interface {
	SetCommonHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	SetHttpClient(client *http.Client) ClientBuilder
	SetUserAgent(agent string) ClientBuilder
	DisableAllTimeouts(disable bool) ClientBuilder
	Build() HttpClient

	BuildMockClient() (HttpClient, MockKeeper)
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) SetCommonHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}

func (c *clientBuilder) DisableAllTimeouts(disable bool) ClientBuilder {
	c.timeoutsDisabled = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(agent string) ClientBuilder {
	c.userAgent = agent
	return c
}

func (c *clientBuilder) Build() HttpClient {
	client := httpClient{
		configurator: clientConfigurator{
			settings: c,
		},
	}
	return &client
}

func (c *clientBuilder) BuildMockClient() (HttpClient, MockKeeper) {
	client := httpClientMock{
		configurator: clientConfigurator{
			settings: c,
		},
		keeper: MockKeeper{
			mocks: make(map[string]*Mock),
		},
	}
	return &client, client.keeper
}
