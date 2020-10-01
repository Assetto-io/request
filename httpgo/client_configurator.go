package httpgo

import (
	"encoding/json"
	"encoding/xml"
	"github.com/assetto-io/request/text"
	"net/http"
	"strings"
)

type clientConfigurator struct {
	settings *clientBuilder
}

func (c *clientConfigurator) mapRequestBody(contentType string, body interface{}) ([]byte, error) {
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

func (c *clientConfigurator) getHeaders(headers ...http.Header) http.Header {
	// we are interested only in one httpgo.Header
	if len(headers) > 0 {
		return headers[0]
	}

	return http.Header{}
}

func (c *clientConfigurator) mapRequestHeaders(customHeaders http.Header) http.Header {
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
