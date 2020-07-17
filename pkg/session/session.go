package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/equipindustry/visanetgo/pkg/visanet"
	"github.com/equipindustry/visanetgo/utils"

	"github.com/equipindustry/visanetgo/pkg/client"
)

// Session struct.
type Session struct {
	Amount              float64   `json:"amount"`
	Antifraud           Antifraud `json:"antifraud"`
	Channel             string    `json:"channel"`
	RecurrenceMaxAmount *float64  `json:"recurrenceMaxAmount"`
}

// Antifraud struct.
type Antifraud struct {
	ClientIP           string             `json:"clientIp"`
	MerchantDefineData MerchantDefineData `json:"merchantDefineData"`
}

// MerchantDefineData struct.
type MerchantDefineData struct {
	Mdd4  string `json:"MDD4"`
	Mdd32 string `json:"MDD32"`
	Mdd75 string `json:"MDD75"`
	Mdd77 int    `json:"MDD77"`
}

// ResponseSession struct.
type ResponseSession struct {
	SessionKey     string
	ExpirationTime int
}

// CreateSession returns struct and string.
func (s *Session) CreateSession(token string) (*ResponseSession, error) {
	url := fmt.Sprintf(utils.GetSessionAPIURL(visanet.Credentials.IsDevelopment), visanet.Credentials.MerchantID)
	sessionResponse := ResponseSession{}

	header := http.Header{
		"Authorization": {
			token,
		},
		"Content-Type": {
			"application/json",
		},
	}

	byteArray, err := json.Marshal(s)
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
