package example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEndpoints(t *testing.T) {
	// Init:

	// Exec:
	endpoints, err := GetEndpoints()

	// Validation:
	assert.NotNil(t, endpoints)
	assert.Nil(t, err)
}
