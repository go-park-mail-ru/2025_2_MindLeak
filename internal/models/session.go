package models

import "github.com/google/uuid"

type Session struct {
	SessionId uuid.UUID
	UserId    uuid.UUID
}
