package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
	Name     string
	Avatar   string
}
