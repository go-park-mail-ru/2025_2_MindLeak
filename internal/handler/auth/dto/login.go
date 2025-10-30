package dto

type UserInputLogin struct {
	Email    string
	Password string
}

type UserOutputLogin struct {
	Name   string
	Avatar string
	Email  string
}
