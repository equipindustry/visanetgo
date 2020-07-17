package authorization

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/equipindustry/visanetgo/pkg/client"
	"github.com/equipindustry/visanetgo/pkg/visanet"
	"github.com/equipindustry/visanetgo/utils"
)

// Authorization struct.
type Authorization struct {
	Antifraud   *string     `json:"antifraud"`
	CaptureType string      `json:"captureType"`
	CardHolder  *CardHolder `json:"cardHolder,omitempty"`
	Channel     string      `json:"channel"`
	Countable   bool        `json:"countable"`
	Order       Order       `json:"order"`
	Recurrence  *Recurrence `json:"recurrence,omitempty"`
	Sponsored   *string     `json:"sponsored"`
}

// CardHolder struct.
type CardHolder struct {
	DocumentNumber string `json:"documentNumber"`
	DocumentType   string `json:"documentType"`
}

// Order struct.
type Order struct {
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	ProductID      *string `json:"productId,omitempty"`
	PurchaseNumber string  `json:"purchaseNumber"`
	TokenID        string  `json:"tokenId"`
}

// Recurrence struct.
type Recurrence struct {
	Amount        float64 `json:"amount"`
	BeneficiaryID string  `json:"beneficiaryId"`
	Frequency     string  `json:"frequency"`
	MaxAmount     float64 `json:"maxAmount"`
	Type          string  `json:"type"`
}

// CreateAuthorization returns struct and string.
func (a *Authorization) CreateAuthorization(token string) (*ResponseAuthorization, error) {
	url := fmt.Sprintf(utils.GetAuthorizationAPIURL(visanet.Credentials.IsDevelopment), visanet.Credentials.MerchantID)
	authorizationResponse := ResponseAuthorization{}

	header := http.Header{
		"Authorization": {
			token,
		},
		"Content-Type": {
			"application/json",
		},
	}

	byteArray, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	body, err := client.Client(http.MethodPost, url, bytes.NewBuffer(byteArray), header)
	if err != nil {
		if body != nil {
			var errResponse ErrorAuthorization
			if err := json.Unmarshal(body, &errResponse); err != nil {
				return nil, err
			}
			return nil, errors.New(errResponse.Data.ActionDescription)
		}
		return nil, err
	}

	if err := json.Unmarshal(body, &authorizationResponse); err != nil {
		return nil, err
	}

	return &authorizationResponse, nil
}
