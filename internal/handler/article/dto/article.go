package dto

import (
	"time"

	"github.com/google/uuid"
)

type ArticleOutputDTO struct {
	Id           uuid.UUID `json:"-"`
	AuthorId     uuid.UUID `json:"-"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"-"`
	Image        string    `json:"image"`
	AuthorName   string    `json:"author_name"`
	AuthorAvatar string    `json:"author_avatar"`
}
