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
	ID            uuid.UUID `db:"article_id"`
	AuthorID      uuid.UUID `db:"author_id"`
	Title         string    `db:"title"`
	Content       string    `db:"content"`
	MediaURL      string    `db:"media_url,omitempty"` // URL из MinIO
	TopicID       uuid.UUID `db:"topic_id"`
	Status        string    `db:"status"` // draft, published, archived
	CommentsCount int       `db:"comments_count"`
	RepostsCount  int       `db:"reposts_count"`
	ViewsCount    int       `db:"views_count"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
	Topic         Topic     `db:"-"`
	AuthorName    string    `db:"-"`
	AuthorAvatar  string    `db:"-"`
}
