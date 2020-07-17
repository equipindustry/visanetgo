package visanet

import "sync"

var (
	doOnce sync.Once
	// Credentials declaration.
	Credentials *credential
)

type credential struct {
	User          string
	Password      string
	MerchantID    string
	IsDevelopment bool
}

func init() {
	doOnce.Do(func() {
		Credentials = &credential{}
	})
}

// CreateCredential instance credential.
func CreateCredential(user, password, merchantID string, isDevelopment bool) {
	Credentials.User = user
	Credentials.Password = password
	Credentials.MerchantID = merchantID
	Credentials.IsDevelopment = isDevelopment
}
