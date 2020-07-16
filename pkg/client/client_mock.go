package client

import (
	"io"
	"net/http"
)

// HTTPClient interface.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// HClient declaration.
var HClient HTTPClient

func init() {
	HClient = &http.Client{}
}

// Post sends a post request to the URL with the body.
func Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return HClient.Do(request)
}
