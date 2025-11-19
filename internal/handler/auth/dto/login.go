package dto

import "github.com/google/uuid"

type UserInputLogin struct {
	Email    string
	Password string
}

type UserOutputLogin struct {
	Id     uuid.UUID
	Name   string
	Avatar string
	Email  string
}
