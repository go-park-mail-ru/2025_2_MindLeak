package models

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	Id        uuid.UUID  `db:"comment_id"`
	ArticleId uuid.UUID  `db:"article_id"`
	UserId    uuid.UUID  `db:"user_id"`
	Content   string     `db:"content"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	ReplyTo   *uuid.UUID `db:"reply_to"`

	AuthorName   string `db:"-"`
	AuthorAvatar string `db:"-"`

	ArticleTitle string `db:"-"`
}
