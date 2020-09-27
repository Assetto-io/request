package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

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

	return client.Do(request)
}

func (c *httpClient) mapRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)

	// get headers from default settings
	for key, value := range c.Headers {
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
