package models

type UserRegister struct {
	FirstName   string
	LastName    string
	DateOfBirth string
	Password    string
	Gender      string
	Email       string
}

type UpdateRequest struct {
	FirstName   string
	LastName    string
	DateOfBirth string
	Gender      string
}
