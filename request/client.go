package request

import (
	"net/http"
	"sync"
)

// HttpClient provides a clean way to perform an HTTP call
type HttpClient interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

type httpClient struct {
	client       *http.Client
	configurator clientConfigurator
	clientOnce   sync.Once
}

func (c *httpClient) Get(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, c.configurator.getHeaders(headers...), nil)
}

func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), body)
}

func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), body)
}

func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), body)
}

func (c *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), nil)
}

func (c *httpClient) Options(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, c.configurator.getHeaders(headers...), nil)
}
