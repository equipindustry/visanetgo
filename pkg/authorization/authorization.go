package authorization

import (
	"bytes"
	"encoding/json"
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

// ResponseAuthorization struct.
type ResponseAuthorization struct {
	Header  Header        `json:"header"`
	Order   OrderResponse `json:"order"`
	DataMap DataMap       `json:"dataMap"`
}

// DataMap struct.
type DataMap struct {
	Currency          string `json:"CURRENCY"`
	TransactionDate   string `json:"TRANSACTION_DATE"`
	Terminal          string `json:"TERMINAL"`
	ActionCode        string `json:"ACTION_CODE"`
	TraceNumber       string `json:"TRACE_NUMBER"`
	EciDescription    string `json:"ECI_DESCRIPTION"`
	Eci               string `json:"ECI"`
	Brand             string `json:"BRAND"`
	Card              string `json:"CARD"`
	Merchant          string `json:"MERCHANT"`
	Status            string `json:"STATUS"`
	Adquirente        string `json:"ADQUIRENTE"`
	ActionDescription string `json:"ACTION_DESCRIPTION"`
	IDUnico           string `json:"ID_UNICO"`
	Amount            string `json:"AMOUNT"`
	ProcessCode       string `json:"PROCESS_CODE"`
	RecurrenceStatus  string `json:"RECURRENCE_STATUS"`
	TransactionID     string `json:"TRANSACTION_ID"`
	AuthorizationCode string `json:"AUTHORIZATION_CODE"`
}

// Header struct.
type Header struct {
	EcoreTransactionUUID string `json:"ecoreTransactionUUID"`
	EcoreTransactionDate int64  `json:"ecoreTransactionDate"`
	Millis               int64  `json:"millis"`
}

// OrderResponse struct.
type OrderResponse struct {
	TokenID           string  `json:"tokenId"`
	PurchaseNumber    string  `json:"purchaseNumber"`
	ProductID         string  `json:"productId"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	AuthorizedAmount  float64 `json:"authorizedAmount"`
	AuthorizationCode string  `json:"authorizationCode"`
	ActionCode        string  `json:"actionCode"`
	TraceNumber       string  `json:"traceNumber"`
	TransactionDate   string  `json:"transactionDate"`
	TransactionID     string  `json:"transactionId"`
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
		return nil, err
	}

	if err := json.Unmarshal(body, &authorizationResponse); err != nil {
		return nil, err
	}

	return &authorizationResponse, nil
}
