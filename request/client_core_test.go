package request

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type testData struct {
	XMLName   xml.Name `xml:"person"`
	Firstname string   `xml:"firstname" json:"first_name"`
	Lastname  string   `xml:"lastname" json:"last_name"`
}

func TestMapRequestHeaders(t *testing.T) {
	// Init
	client := httpClient{}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", "awesome-http-client")
	client.SetCommonHeaders(headers)

	customHeaders := make(http.Header)
	customHeaders.Set("X-Request-id", "assetto-321")

	// Exec
	finalHeaders := client.mapRequestHeaders(customHeaders)

	// Validation
	assert.Equal(t, 3, len(finalHeaders))
	assert.Equal(t, "assetto-321", finalHeaders.Get("X-Request-id"))
	assert.Equal(t, "awesome-http-client", finalHeaders.Get("User-Agent"))
	assert.Equal(t, "application/json", finalHeaders.Get("Content-Type"))
}

func TestMapRequestBody(t *testing.T) {
	// Init
	client := httpClient{}
	data := testData{
		XMLName:   xml.Name{Local: "Person"},
		Firstname: "Denis",
		Lastname:  "Denisov",
	}

	t.Run("RequestBodyIsNil", func(t *testing.T) {
		// Exec
		body, err := client.mapRequestBody("", nil)

		// Validation
		assert.Nil(t, body)
		assert.Nil(t, err)
	})

	t.Run("RequestBodyDefault", func(t *testing.T) {
		// Exec
		body, err := client.mapRequestBody("", data)
		buffer := testData{}
		json.Unmarshal(body, &buffer)

		// Validation
		assert.NotNil(t, body)
		assert.Nil(t, err)
		assert.Equal(t, data.Lastname, buffer.Lastname)
		assert.Equal(t, "Denis", buffer.Firstname)
	})

	t.Run("RequestBodyIsJson", func(t *testing.T) {
		// Exec
		body, err := client.mapRequestBody("application/json", data)
		buffer := testData{}
		json.Unmarshal(body, &buffer)

		// Validation
		assert.NotNil(t, body)
		assert.Nil(t, err)
		assert.Equal(t, data.Lastname, buffer.Lastname)
		assert.Equal(t, "Denis", buffer.Firstname)
	})

	t.Run("RequestBodyIsXml", func(t *testing.T) {
		// Exec
		body, err := client.mapRequestBody("application/xml", data)
		buffer := testData{}
		xml.Unmarshal(body, &buffer)

		// Validation
		assert.NotNil(t, body)
		assert.Nil(t, err)
		assert.Equal(t, data.Lastname, buffer.Lastname)
		assert.Equal(t, "Denis", buffer.Firstname)
	})

}
