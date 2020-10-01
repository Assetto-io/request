package httpgo

import (
	"bytes"
	"errors"
	"github.com/assetto-io/request/text"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	defaultIdleConnection    = 5
	defaultConnectionTimeout = 5 * time.Second
	defaultResponseTimeout   = 5 * time.Second
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.configurator.mapRequestHeaders(headers)

	requestBody, err := c.configurator.mapRequestBody(fullHeaders.Get(text.HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new httpgo")
	}

	request.Header = fullHeaders

	client := c.getClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resultResponse := Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
	}

	return &resultResponse, nil
}

func (c *httpClient) getClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.configurator.settings.client != nil {
			c.client = c.configurator.settings.client
			return
		}

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
	if c.configurator.settings.maxIdleConnections > 0 {
		return c.configurator.settings.maxIdleConnections
	}
	return defaultIdleConnection
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.configurator.settings.timeoutsDisabled {
		return 0
	}

	if c.configurator.settings.connectionTimeout > 0 {
		return c.configurator.settings.connectionTimeout
	}
	return defaultConnectionTimeout
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.configurator.settings.timeoutsDisabled {
		return 0
	}

	if c.configurator.settings.responseTimeout > 0 {
		return c.configurator.settings.responseTimeout
	}
	return defaultResponseTimeout
}
