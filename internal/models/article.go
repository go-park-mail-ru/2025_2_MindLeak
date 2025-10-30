package models

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	Id           uuid.UUID
	AuthorId     uuid.UUID
	Title        string
	Content      string
	CreatedAt    time.Time
	Image        string
	AuthorName   string
	AuthorAvatar string
}
