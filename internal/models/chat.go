package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID `json:"id" db:"room_id"`
	Name      string    `json:"name" db:"name"`
	IsGroup   bool      `json:"is_group" db:"is_group"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Message struct {
	ID        uuid.UUID `json:"id" db:"message_id"`
	RoomID    uuid.UUID `json:"room_id" db:"room_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	UserName  string    `json:"user_name" db:"user_name"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Text      string    `json:"text" db:"text"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
