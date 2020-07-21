package session

import (
	"encoding/json"
	"fmt"
)

// ISessionResponse interface.
type ISessionResponse interface {
	SetSession(session string)
	GetSession() *Response
}

// ISessionFailResponse interface.
type ISessionFailResponse interface {
	SetFailSession(session string)
	GetFailSession() *MessageError
}

// NewSessionSuccessResponseFactory return ISessionResponse.
func NewSessionSuccessResponseFactory() ISessionResponse {
	return &Response{}
}

// NewSessionFailResponseFactory return ISessionFailResponse.
func NewSessionFailResponseFactory() ISessionFailResponse {
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
		fmt.Println(err)
	}
}

// GetSession return Response.
func (s *Response) GetSession() *Response {
	return s
}

// SetFailSession func.
func (m *MessageError) SetFailSession(session string) {
	if err := json.Unmarshal([]byte(session), &m); err != nil {
		fmt.Println(err)
	}
}

// GetFailSession return MessageError.
func (m *MessageError) GetFailSession() *MessageError {
	return m
}
