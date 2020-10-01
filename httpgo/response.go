package httpgo

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJSON(target interface{}) error {
	return json.Unmarshal(r.Body, target)
}

func (r *Response) UnmarshalXML(target interface{}) error {
	return xml.Unmarshal(r.Body, target)
}
