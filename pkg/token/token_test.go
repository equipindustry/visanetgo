package token

import (
	"fmt"
	"testing"

	"github.com/equipindustry/visanetgo/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateTokenSuccess(t *testing.T) {
	tokenFactory := NewTokenFactory()
	tokenFactory.SetToken(test.TokenResponse)
	assert.EqualValues(t, tokenFactory.GetToken(), test.TokenResponse)
}

func TestCreateTokenFail(t *testing.T) {
	tokenFactory := NewTokenFactory()
	tokenFactory.SetToken(test.TokenFailResponse)
	fmt.Println(tokenFactory.GetToken())
	assert.EqualValues(t, tokenFactory.GetToken(), test.TokenFailResponse, test.TokenFailResponse)
}
