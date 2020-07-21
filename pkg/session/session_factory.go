package session

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Responser interface.
type Responser interface {
	SetSession(session string)
	GetSession() *Response
}

// FailResponser interface.
type FailResponser interface {
	SetFailSession(session string)
	GetFailSession() *MessageError
}

// NewSessionSuccessResponseFactory return Responser.
func NewSessionSuccessResponseFactory() Responser {
	return &Response{}
}

// NewSessionFailResponseFactory return FailResponser.
func NewSessionFailResponseFactory() FailResponser {
	return &MessageError{}
}

// MessageError struct.
type MessageError struct {
	ErrorCode    int64  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Data         data   `json:"data"`
}

// data struct.
type data struct {
}

// Response struct.
type Response struct {
	SessionKey     string
	ExpirationTime int
}

// SetSession func.
func (s *Response) SetSession(session string) {
	if err := json.Unmarshal([]byte(session), &s); err != nil {
		log.Debug(err)
	}
}

// GetSession return Response.
func (s *Response) GetSession() *Response {
	return s
}

// SetFailSession func.
func (m *MessageError) SetFailSession(session string) {
	if err := json.Unmarshal([]byte(session), &m); err != nil {
		log.Debug(err)
	}
}

// GetFailSession return MessageError.
func (m *MessageError) GetFailSession() *MessageError {
	return m
}
