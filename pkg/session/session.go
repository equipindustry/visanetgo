package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/equipindustry/visanetgo/model"
	"github.com/equipindustry/visanetgo/pkg/client"
	"github.com/spf13/viper"
)

// CreateSession returns struct and string.
func CreateSession(token string, sessionBody model.SessionBody) (*model.SessionResponse, error) {
	url := fmt.Sprintf("%s/%s", viper.GetString("VISANET_DEV_SESSION_API"), viper.GetString("VISANET_DEV_MERCHANT_ID"))
	sessionResponse := model.SessionResponse{}

	header := http.Header{
		"Authorization": {
			token,
		},
		"Content-Type": {
			"application/json",
		},
	}

	byteArray, err := json.Marshal(sessionBody)
	if err != nil {
		fmt.Println(err)
	}

	body, err := client.Client(http.MethodPost, url, bytes.NewBuffer(byteArray), header)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &sessionResponse); err != nil {
		return nil, err
	}

	return &sessionResponse, nil
}
