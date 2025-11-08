package dto

import (
	"github.com/google/uuid"
	"time"
)

type CommentDto struct {
	Id        uuid.UUID
	ArticleId uuid.UUID
	UserId    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	ReplyTo   *uuid.UUID

	AuthorName   string
	AuthorAvatar string

	ArticleTitle string
}
