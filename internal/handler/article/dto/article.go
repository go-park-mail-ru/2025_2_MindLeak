package dto

import (
	"time"

	"github.com/google/uuid"
)

type ArticleOutputDTO struct {
	Id           uuid.UUID `json:"-"`
	AuthorId     uuid.UUID `json:"author_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"-"`
	TopicTitle   string    `json:"topic_title"`
	AuthorName   string    `json:"author_name"`
	AuthorAvatar string    `json:"author_avatar"`
}
