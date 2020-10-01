package httpgo

import (
	"net/http"
)

// Mock structure provides a clean way to configure HTTP mocks based on
// the combination between HTTP method, URL and HTTP body (md5 hash).
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseStatusCode int
	ResponseBody       string
	ResponseHeaders    http.Header
}
