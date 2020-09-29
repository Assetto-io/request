package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	defaultIdleConnection    = 5
	defaultConnectionTimeout = 5 * time.Second
	defaultResponseTimeout   = 5 * time.Second
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	fullHeaders := c.mapRequestHeaders(headers)

	requestBody, err := c.mapRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	request.Header = fullHeaders

	client := c.getClient()

	return client.Do(request)
}

func (c *httpClient) mapRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)

	// get headers from default settings
	for key, value := range c.settings.headers {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}

	// get headers from custom settings
	for key, value := range customHeaders {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}

	return result
}

func (c *httpClient) getClient() *http.Client {
	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.settings.maxIdleConnections > 0 {
		return c.settings.maxIdleConnections
	}
	return defaultIdleConnection
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.settings.timeoutsDisabled {
		return 0
	}

	if c.settings.connectionTimeout > 0 {
		return c.settings.connectionTimeout
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.settings.timeoutsDisabled {
		return 0
	}

	if c.settings.responseTimeout > 0 {
		return c.settings.responseTimeout
	}
	return defaultResponseTimeout
}

func (c *httpClient) mapRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
