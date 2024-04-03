package models

type User struct {
	//gorm.Model

	ID    int64
	Token string

	Email          string
	Username       string
	HashedPassword string
}

type RegistrationRequest struct {
	Username string `json:"username" validate:"required,username" binding:"required"`
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Password string `json:"password" validate:"required,password" binding:"required"`
}

type RegistrationResponse struct {
	Token     string `json:"token"`
	TimeStamp int64  `json:"stamp"`
}

type AuthenticationResponse struct {
	Token string `json:"token"`
}

type AuthenticationRequest struct {
	RegistrationRequest
}
