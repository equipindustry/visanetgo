package model

// SessionBody struct.
type SessionBody struct {
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

// SessionResponse struct.
type SessionResponse struct {
	SessionKey     string `json:"sessionKey"`
	ExpirationTime int    `json:"expirationTime"`
}
