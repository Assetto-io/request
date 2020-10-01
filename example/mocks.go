package example

import (
	"github.com/assetto-io/request/request"
	"net/http"
)

// Put all your Mock settings here
func setMocks() {
	mockError := request.Mock{
		Method:             http.MethodGet,
		Url:                "https://api.github.com",
		ResponseStatusCode: http.StatusOK,
		ResponseBody:       `{"hub_url": "https://api.github.com/hub"}`,
	}
	mockKeeper.AddMock(mockError)
}
