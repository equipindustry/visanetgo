package session

import (
	"fmt"
	"testing"

	"github.com/equipindustry/visanetgo/pkg/api/status"

	"github.com/equipindustry/visanetgo/test"
	"github.com/stretchr/testify/assert"
)

func TestSession_CreateSession(t *testing.T) {
	sessionFactory := NewSessionSuccessResponseFactory()
	sessionFactory.SetSession(test.SessionResponse)
	assert.NotNil(t, sessionFactory.GetSession())
}

func TestSession_CreateFailTokenSession(t *testing.T) {
	sessionFactory := NewSessionFailResponseFactory()
	sessionFactory.SetFailSession(test.SessionResponseFailTokenUsed)
	assert.EqualValues(t, sessionFactory.GetFailSession().ErrorCode, status.BadRequest)
	assert.EqualValues(t, sessionFactory.GetFailSession().ErrorMessage, "Token has been used before")
}

func TestSession_CreateFailDataSession(t *testing.T) {
	sessionFactory := NewSessionFailResponseFactory()
	body := fmt.Sprintf(test.SessionResponseFailSendData, "AMOUNT")
	sessionFactory.SetFailSession(body)
	assert.EqualValues(t, sessionFactory.GetFailSession().ErrorCode, status.BadRequest)
	assert.EqualValues(t, sessionFactory.GetFailSession().ErrorMessage, "AMOUNT is not valid.")
}
