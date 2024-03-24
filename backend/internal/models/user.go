package models

type User struct {
	//gorm.Model

	ID             uint32
	Email          string
	Username       string
	HashedPassword string
}

type RegistrationRequest struct {
	Username string `json:"email" validate:"required,username" binding:"required"`
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Password string `json:"password" validate:"required,password" binding:"required"`
}

type AuthenticationRequest struct {
	RegistrationRequest
}
