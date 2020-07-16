package mock

import "net/http"

// Client is the mock client.
type Client struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// GetDoFunc fetches the mock client's `Do` func.
var GetDoFunc func(req *http.Request) (*http.Response, error)

// Do is the mock client's `Do` func/
func (m *Client) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
