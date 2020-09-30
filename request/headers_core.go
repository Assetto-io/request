package request

import (
	"github.com/assetto-io/request/text"
	"net/http"
)

func getHeaders(headers ...http.Header) http.Header {
	// we are interested only in one http.Header
	if len(headers) > 0 {
		return headers[0]
	}

	return http.Header{}
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

	// set User-Agent if it is empty
	if c.settings.userAgent != "" {
		if result.Get(text.HeaderUserAgent) != "" {
			return result
		}
		result.Set(text.HeaderUserAgent, c.settings.userAgent)
	}

	return result
}
