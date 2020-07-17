package token

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/equipindustry/visanetgo/pkg/visanet"
	"github.com/equipindustry/visanetgo/utils"

	"github.com/equipindustry/visanetgo/pkg/client"
)

// CreateToken returns a string
func CreateToken() (*string, error) {
	url := utils.GetSecurityAPIURL(visanet.Credentials.IsDevelopment)
	credential := fmt.Sprintf("%s:%s", visanet.Credentials.User, visanet.Credentials.Password)
	encodedCredential := base64.StdEncoding.EncodeToString([]byte(credential))

	header := http.Header{
		"Authorization": {
			fmt.Sprintf("Basic %s", encodedCredential),
		},
	}

	body, err := client.Client(http.MethodPost, url, nil, header)
	if err != nil {
		return nil, err
	}

	token := string(body)

	return &token, nil
}
