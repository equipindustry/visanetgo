package client

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/equipindustry/visanetgo/pkg/api/status"
)

// Client returns a response byte.
func Client(method string, url string, body io.Reader, headers http.Header) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = headers

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	obj, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= status.Ok && res.StatusCode <= status.Accepted {
		return obj, nil
	}

	return obj, err
}
