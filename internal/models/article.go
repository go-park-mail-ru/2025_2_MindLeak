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
	ImageURL     string        `db:"image_url"`
	Status       ArticleStatus `db:"status"`
	PublishedAt  time.Time     `db:"published_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	AuthorName   string        `db:"-"`
	AuthorAvatar string        `db:"-"`
}
