package dto

import "github.com/google/uuid"

type RegisteredUserDto struct {
	Id     uuid.UUID
	Email  string
	Name   string
	Avatar string
}
