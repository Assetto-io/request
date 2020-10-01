package request

import "net/http"

type httpClientMock struct {
	keeper       MockKeeper
	configurator clientConfigurator
}

func (c *httpClientMock) Get(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, c.configurator.getHeaders(headers...), nil)
}

func (c *httpClientMock) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), body)
}

func (c *httpClientMock) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), body)
}

func (c *httpClientMock) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), body)
}

func (c *httpClientMock) Delete(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, c.configurator.getHeaders(headers...), nil)
}

func (c *httpClientMock) Options(url string, headers ...http.Header) (*Response, error) {
	return c.do(http.MethodOptions, url, c.configurator.getHeaders(headers...), nil)
}
