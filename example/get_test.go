package example

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	// You can use ENV Variables or other configuration settings
	// Just for easy example
	useMock = true
)

func TestMain(m *testing.M) {
	// Set you HTTP Client depending on your configuration

	if useMock {
		mockKeeper.DeleteMocks()
		setMockHttpClient()
	} else {
		setHttpClient()
	}

	code := m.Run()
	os.Exit(code)
}

func TestGetEndpoints(t *testing.T) {
	// Init:

	// Exec:
	endpoints, err := GetEndpoints()

	// Validation:
	assert.NotNil(t, endpoints)
	assert.Nil(t, err)
	assert.Equal(t, "https://api.github.com/hub", endpoints.HubUrl)
}
