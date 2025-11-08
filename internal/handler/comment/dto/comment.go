package dto

import (
	"github.com/google/uuid"
	"time"
)

type CommentIODto struct { // Comment Input/Output DTO
	Id        uuid.UUID  `json:"id"`
	ArticleId uuid.UUID  `json:"article_id"`
	UserId    uuid.UUID  `json:"user_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ReplyTo   *uuid.UUID `json:"reply_to"`

	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`

	ArticleTitle string `json:"article_title"`
}

type CommentsOutputDto struct {
	Comments []CommentIODto `json:"comments"`
}
