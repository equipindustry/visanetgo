package authorization

// ResponseAuthorization struct.
type ResponseAuthorization struct {
	Header      Header        `json:"header"`
	Fulfillment Fulfillment   `json:"fulfillment"`
	Order       OrderResponse `json:"order"`
	DataMap     DataMap       `json:"dataMap"`
}

// Fulfillment struct.
type Fulfillment struct {
	Channel     string `json:"channel"`
	MerchantID  string `json:"merchantId"`
	TerminalID  string `json:"terminalId"`
	CaptureType string `json:"captureType"`
	Countable   bool   `json:"countable"`
	FastPayment bool   `json:"fastPayment"`
	Signature   string `json:"signature"`
}

// Header struct.
type Header struct {
	EcoreTransactionUUID string `json:"ecoreTransactionUUID"`
	EcoreTransactionDate int64  `json:"ecoreTransactionDate"`
	Millis               int64  `json:"millis"`
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

// ErrorAuthorization struct.
type ErrorAuthorization struct {
	ErrorCode    int64       `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Header       HeaderError `json:"header"`
	Data         Data        `json:"data"`
}

// Data struct.
type Data struct {
	Currency          string `json:"CURRENCY"`
	TransactionDate   string `json:"TRANSACTION_DATE"`
	Terminal          string `json:"TERMINAL"`
	ActionCode        string `json:"ACTION_CODE"`
	TraceNumber       string `json:"TRACE_NUMBER"`
	EciDescription    string `json:"ECI_DESCRIPTION"`
	Eci               string `json:"ECI"`
	Card              string `json:"CARD"`
	Brand             string `json:"BRAND"`
	Merchant          string `json:"MERCHANT"`
	Status            string `json:"STATUS"`
	Adquirente        string `json:"ADQUIRENTE"`
	ActionDescription string `json:"ACTION_DESCRIPTION"`
	IDUnico           string `json:"ID_UNICO"`
	Amount            string `json:"AMOUNT"`
	ProcessCode       string `json:"PROCESS_CODE"`
	TransactionID     string `json:"TRANSACTION_ID"`
}

// HeaderError struct.
type HeaderError struct {
	EcoreTransactionUUID string `json:"ecoreTransactionUUID"`
	EcoreTransactionDate int64  `json:"ecoreTransactionDate"`
	Millis               int64  `json:"millis"`
}
