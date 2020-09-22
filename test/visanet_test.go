package test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/equipindustry/visanetgo/pkg/api/status"

	"github.com/equipindustry/visanetgo/mock"
	"github.com/equipindustry/visanetgo/pkg/client"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	client.HClient = &mock.Client{}
}

func TestCreateToken(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(TokenResponse)))

	mock.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: status.Created,
			Body:       r,
		}, nil
	}

	user := viper.Get("VISANET_DEV_USER")
	pwd := viper.Get("VISANET_DEV_PASSWD")
	url := viper.GetString("VISANET_DEV_SECURITY_API")
	encodedCredential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pwd)))
	fmt.Println(encodedCredential)

	header := http.Header{
		"Authorization": {
			"Basic " + encodedCredential,
		},
	}

	response, err := client.Post(url, nil, header)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.EqualValues(t, TokenResponse, string(body))
}

func TestCreateSession(t *testing.T) {
	url := fmt.Sprintf("%s/%s", viper.GetString("VISANET_DEV_SESSION_API"), viper.GetString("VISANET_DEV_MERCHANT_ID"))
	r := ioutil.NopCloser(bytes.NewReader([]byte(SessionResponse)))

	mock.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: status.Ok,
			Body:       r,
		}, nil
	}

	header := http.Header{
		"Authorization": {
			TokenResponse,
		},
	}

	response, err := client.Post(url, nil, header)
	if err != nil {
		t.Error(err)
	}

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.EqualValues(t, "e242ccbf2e4ab1bbd705e20cbebf9bd3df739e7a0239a14294b300d83995b092", "e242ccbf2e4ab1bbd705e20cbebf9bd3df739e7a0239a14294b300d83995b092")
}

func TestCreateAuthorization(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(AuthorizationResponse)))

	mock.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: status.Created,
			Body:       r,
		}, nil
	}

	header := http.Header{
		"Authorization": {
			TokenResponse,
		},
	}

	url := viper.GetString("VISANET_DEV_AUTHORIZATION_API")

	response, err := client.Post(url, nil, header)
	if err != nil {
		t.Error(err)
	}

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.EqualValues(t, "44BDC4D1500F4927BDC4D1500F7927D6", "44BDC4D1500F4927BDC4D1500F7927D6")
}
