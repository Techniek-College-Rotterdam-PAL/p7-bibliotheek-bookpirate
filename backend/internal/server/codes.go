package server

import (
	"encoding/json"
)

type Code uint32

const (
	TCRStudentDomain    string = "student.zadkine.nl"
	TCRDocentDomain     string = "tcrmbo.nl"
	authorizationHeader string = "authorization"
	defaultAuthLength   int    = 32
)

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
	InsufficientPermissions
	UsernameAlreadyTaken
	EmailAlreadyUsed
	TooManyBooks
	AlreadyLoggedIn
	AlreadyLoggedInDifferentAccount
	InvalidEmail
	AdminNeeded
	InvalidPassword
	InvalidAuthentication
	IsbnAlreadyFound
	IsbnNotFound
	SuccessfulAuthentication
	SuccessfulDeAuthentication
	SuccessfulRegistration
	SuccessfulInsert
	SuccessfulReservation
	UnsuccessfulRegistration
	InvalidSession
	NoMoreStock
)

var messages = map[Code]string{
	MethodNotAllowed:    "Method is not allowed that this endpoint.",
	MalformedContent:    "Malformed body or invalid content.",
	IncorrectPassword:   "Incorrect password.",
	InternalServerError: "Internal server error.",
	ClientRateLimit:     "Temporarily blocked",
	Forbidden:           "Forbidden.",
	NotFound:            "Not Found",

	UserNotFound:                    "User not found.",
	InvalidAuthenticationRequest:    "Malformed body or invalid content.",
	DatabaseQueryError:              "Internal server error.",
	InsufficientPermissions:         "User does not have permission",
	UsernameAlreadyTaken:            "Username already in use",
	IsbnAlreadyFound:                "Book with ISBN already added",
	IsbnNotFound:                    "Unknown Book",
	InvalidEmail:                    "Invalid Email Domain",
	AdminNeeded:                     "Admin Needed",
	InvalidPassword:                 "Invalid Password",
	EmailAlreadyUsed:                "Email already used",
	TooManyBooks:                    "Book already reserved",
	AlreadyLoggedIn:                 "Already logged in",
	AlreadyLoggedInDifferentAccount: "Already logged into different account, please logout to login.html into different account",
	InvalidAuthentication:           "invalid email or password",
	SuccessfulAuthentication:        "Successfully authenticated.",
	SuccessfulDeAuthentication:      "Successfully de-authenticated.",
	SuccessfulRegistration:          "Successfully registered.",
	SuccessfulInsert:                "Successfully added book.",
	SuccessfulReservation:           "Successfully reserved book",
	UnsuccessfulRegistration:        "Internal Server Error.",
	InvalidSession:                  "Invalid session.",
	NoMoreStock:                     "No stock left",
}

type Message struct {
	Data    json.RawMessage `json:"data,omitempty"`
	Code    uint32          `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
}
