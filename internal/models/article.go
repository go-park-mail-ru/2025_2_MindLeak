package models

import (
	"time"

	"github.com/google/uuid"
)

type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "draft"
	ArticleStatusPublished ArticleStatus = "published"
	ArticleStatusArchived  ArticleStatus = "archived"
)

type Article struct {
	ID           uuid.UUID     `db:"article_id"`
	AuthorID     uuid.UUID     `db:"author_id"`
	Title        string        `db:"title"`
	Content      string        `db:"content"`
	Status       ArticleStatus `db:"status"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	AuthorName   string        `db:"-"`
	AuthorAvatar string        `db:"-"`

	Topic Topic `db:"-"`
}
