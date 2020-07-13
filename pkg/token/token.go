package token

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/equipindustry/visanetgo/pkg/client"
	"github.com/spf13/viper"
)

// CreateToken returns a string
func CreateToken() string {
	url := viper.GetString("VISANET_DEV_AUTHORIZATION_API")
	user := viper.GetString("VISANET_DEV_USER")
	pwd := viper.GetString("VISANET_DEV_PASSWD")
	encodedCredential := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pwd)))

	header := http.Header{
		"Authorization": {
			fmt.Sprintf("Basic %s", encodedCredential),
		},
	}

	body, err := client.Client(http.MethodPost, url, nil, header)
	if err != nil {
		fmt.Println(err)
	}

	token := string(body)

	return token
}
