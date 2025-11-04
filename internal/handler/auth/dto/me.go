package dto

import "github.com/google/uuid"

type UserOutputMe struct {
	Id     uuid.UUID `json:"id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Avatar string    `json:"avatar"`
}
