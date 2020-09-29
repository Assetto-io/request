package request

import (
	"net/http"
	"time"
)

type ClientBuilder interface {
	SetCommonHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(connections int) ClientBuilder
	DisableAllTimeouts(disable bool) ClientBuilder
	Build() HttpClient
}

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	timeoutsDisabled   bool
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() HttpClient {
	client := httpClient{
		settings: c,
	}
	return &client
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
