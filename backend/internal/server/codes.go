package server

import (
	"encoding/json"
)

type Code uint32

const (
	MethodNotAllowed = iota
	MalformedContent
	IncorrectPassword
	InternalServerError
	ClientRateLimit
	Forbidden
	NotFound

	UserNotFound
	InvalidAuthenticationRequest
	DatabaseQueryError
	UsernameAlreadyTaken
	SuccessfulAuthentication
	SuccessfulDeauthentication
	SuccessfulRegistration
	UnsuccessfulRegistration
	InvalidSession
)

var messages = map[Code]string{
	MethodNotAllowed:    "Method is not allowed that this endpoint.",
	MalformedContent:    "Malformed body or invalid content.",
	IncorrectPassword:   "Incorrect password.",
	InternalServerError: "Internal server error.",
	ClientRateLimit:     "Temporarily blocked",
	Forbidden:           "Forbidden.",
	NotFound:            "Not Found",

	UserNotFound:                 "User not found.",
	InvalidAuthenticationRequest: "Malformed body or invalid content.",
	DatabaseQueryError:           "Internal server error.",
	UsernameAlreadyTaken:         "Username already in use",
	SuccessfulAuthentication:     "Successfully authenticated.",
	SuccessfulDeauthentication:   "Successfully deauthenticated.",
	SuccessfulRegistration:       "Successfully registered.",
	UnsuccessfulRegistration:     "Internal Server Error.",
	InvalidSession:               "Invalid session.",
}

type Message struct {
	Data    json.RawMessage `json:"data,omitempty"`
	Code    uint32          `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
}
