package token

// Tokenizer interface.
type Tokenizer interface {
	SetToken(token string)
	GetToken() string
}

// NewTokenFactory return Tokenizer.
func NewTokenFactory() Tokenizer {
	return &Token{}
}

// Token struct.
type Token struct {
	token string
}

// SetToken func.
func (t *Token) SetToken(token string) {
	t.token = token
}

// GetToken return string.
func (t *Token) GetToken() string {
	return t.token
}
