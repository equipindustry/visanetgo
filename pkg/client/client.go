package client

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/equipindustry/visanetgo/pkg/api/status"
)

var (
	errInvalidRequest = errors.New("The request has invalid syntax")
	errAuthentication = errors.New("The request could not be processed due to problems with the keys")
	errUnexpected     = errors.New("Unexpected error, response code is not handled")
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

	switch res.StatusCode {
	case status.Unauthorized:
		err = errAuthentication
	case status.BadRequest:
		err = errInvalidRequest
	}

	if err != nil {
		err = fmt.Errorf("%v: %s", err, string(obj))
		return obj, err
	}

	if res.StatusCode >= status.Ok && res.StatusCode <= status.Accepted {
		return obj, nil
	}

	return nil, errUnexpected
}
