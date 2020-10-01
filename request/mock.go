package request

import (
	"net/http"
)

// Mock structure provides a clean way to configure HTTP mocks based on
// the combination between request method, URL and request body (md5 hash).
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseStatusCode int
	ResponseBody       string
	ResponseHeaders    http.Header
}
