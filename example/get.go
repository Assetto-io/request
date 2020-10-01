package example

import "fmt"

const (
	URL = "https://api.github.com"
)

type Endpoints struct {
	EventsUrl string `json:"events_url"`
	FeedsUrl  string `json:"feeds_url"`
	HubUrl    string `json:"hub_url"`
}

func GetEndpoints() (*Endpoints, error) {
	// Get httpgo:

	response, err := httpClient.Get(URL, nil)
	if err != nil {
		return nil, err
	}

	// No need for closing httpgo.Response.Body
	// Simple interface for response data extraction:

	fmt.Println(fmt.Sprintf("Status code: %d:", response.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s:", response.Status))
	fmt.Println(fmt.Sprintf("Response body: %s:", response.String()))

	// Simple interface for unmarshaling response data:

	var endpoints Endpoints
	if err := response.UnmarshalJSON(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Unmarshaled feeds_url data: %s:", endpoints.FeedsUrl))
	return &endpoints, nil
}
