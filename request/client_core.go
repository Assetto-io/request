package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/assetto-io/request/text"
	"io/ioutil"
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

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.mapRequestHeaders(headers)

	requestBody, err := c.mapRequestBody(fullHeaders.Get(text.HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create new request")
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
		if c.settings.client != nil {
			c.client = c.settings.client
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
	case text.ContentTypeJson:
		return json.Marshal(body)
	case text.ContentTypeXml:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
