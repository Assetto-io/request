package request

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Response struct {
	status     string
	statusCode int
	headers    http.Header
	body       []byte
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Headers() http.Header {
	return r.headers
}

func (r *Response) Body() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) UnmarshalJSON(target interface{}) error {
	return json.Unmarshal(r.body, target)
}

func (r *Response) UnmarshalXML(target interface{}) error {
	return xml.Unmarshal(r.body, target)
}
