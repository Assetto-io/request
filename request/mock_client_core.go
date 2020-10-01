package request

import (
	"errors"
	"fmt"
	"github.com/assetto-io/request/text"
	"net/http"
)

func (c *httpClientMock) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.configurator.mapRequestHeaders(headers)

	requestBody, err := c.configurator.mapRequestBody(fullHeaders.Get(text.HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	mock := c.keeper.mocks[c.keeper.getMockKey(method, url, string(requestBody))]

	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		resultResponse := Response{
			Status:     c.getResponseStatus(mock.ResponseStatusCode),
			StatusCode: mock.ResponseStatusCode,
			Headers:    mock.ResponseHeaders,
			Body:       []byte(mock.ResponseBody),
		}
		return &resultResponse, nil
	}

	return nil, errors.New(fmt.Sprintf("no mock matching %s from '%s' with given body", method, url))
}

func (c *httpClientMock) getResponseStatus(statusCode int) string {
	return http.Response{
		StatusCode: statusCode,
	}.Status
}
