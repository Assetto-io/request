# request
A developer-friendly &  production-ready HTTP request library for Gopher.

## Installation

```bash
go get github.com/assetto-io/request
```

## Usage
In order to use the library you need to import the corresponding package:

```go
import "github.com/assetto-io/request/httpgo"
```

## Configuration
Once you have imported the package, you can now start using the client. First you need to configure and build the client as you need:

```go
headers := make(http.Header)
headers.Set("Some-Common-Header-If-You-Need-It", "value-for-all-requests")

// Create a new builder:
client := httpgo.NewBuilder().

    // Add as many configurations as you need

	// Configure global headers to be used in every httpgo made by this client:
	SetHeaders(headers).

	// Configure the timeout for getting a new connection:
	SetConnectionTimeout(5 * time.Second).

	// Configure the timeout for performing the actual HTTP call:
	SetResponseTimeout(30 * time.Millisecond).

	// Configure the User-Agent header that will be used for all of the requests:
	SetUserAgent("Your-User-Agent").

	// Finally, build the client and start using it!
	Build()
```

## Performing HTTP calls
The ``HttpClient`` interface provides convenient methods that you can use to perform different HTTP calls.

### Get

```go
const (
    URL = "http://some-custom-url.com"
)

type User struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func GetUser() (*User, error) {
    // Get response

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

	var user User
	if err := response.UnmarshalJSON(&user); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Unmarshaled user data: %s:", user.UserName))
	return &user, nil
}
```

### Post

```go

const (
 URL = "https://api.github.com/user/repos" 
)


// The struct representing the actual JSON response from the API we're calling:
type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

// The struct representing the JSON body we're going to send:
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private"`
}

func CreateRepo(request Repository) (*Repository, error) {
	// Make the httpgo and wait for the response:
	response, err := httpClient.Post(URL, request)
	if err != nil {
		return nil, err
	}

	// Deal with failed status codes:
	if response.StatusCode != http.StatusCreated {
		var githubError GithubError
		if err := response.UnmarshalJson(&githubError); err != nil {
			return nil, errors.New("error processing github error response when creating a new repo")
		}
		return nil, errors.New(githubError.Message)
	}

	// Deal with successful response:
	var result Repository
	if err := response.UnmarshalJson(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
```

## Testing

The library provides mocking requests and getting a particular response.
 The mock key is generated using the ``HTTP method``, the ``request URL`` and the ``request body``. Every request with these same elements will return the same mock.

### Build mock HTTP Client:
```go
headers := make(http.Header)
headers.Set("Some-Common-Header-If-You-Need-It", "value-for-all-requests")

// Create a new builder:
client, mockStorage := httpgo.NewBuilder().

    // Add as many configurations as you need

	// Configure global headers to be used in every httpgo made by this client:
	SetHeaders(headers).

	// Configure the User-Agent header that will be used for all of the requests:
	SetUserAgent("Your-User-Agent").

	// Finally, build the mock-client and start using it!
	BuildMockClient()
```

Once you built mock-client & mock-storage, every request will be handled by this storage and will not be sent against the real API. If there is no mock matching the current request you'll get an error saying ``no mock matching {METHOD} from '{URL}' with given body``.

### Configuring a given HTTP mock:

```go
// Delete all mocks in every new test case to ensure a clean environment:
mockStorage.DeleteMocks()

// Configure a new mock:
mockStorage.AddMock(httpgo.Mock{
	Method:      http.MethodPost,
	Url:         "https://api.github.com/user/repos",
	RequestBody: `{"name":"test-repo","private":true}`,

	Error: errors.New("timeout from github"),
})
```

In this case, we're telling the client that when we send a POST request against that URL and with that body, we want that particular error. In this case, no response was returned. Let's see how you can configure a particular response:


```go
// Delete all mocks in every new test case to ensure a clean environment:
mockStorage.DeleteMocks()

// Configure a new mock:
mockStorage.AddMock(httpgo.Mock{
	Method:      http.MethodPost,
	Url:         "https://api.github.com/user/repos",
	RequestBody: `{"name":"test-repo","private":true}`,

	ResponseStatusCode: http.StatusCreated,
	ResponseBody:       `{"id":123,"name":"test-repo"}`,
})
```

In this case, we get a response with status code ``201 Created`` and that particular response body.
