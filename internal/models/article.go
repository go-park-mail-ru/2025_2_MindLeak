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
	ID            uuid.UUID `json:"id"`
	AuthorID      uuid.UUID `json:"author_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	MediaURL      string    `json:"media_url,omitempty"` // URL из MinIO
	TopicID       uuid.UUID `json:"topic_id"`
	Status        string    `json:"status"` // draft, published, archived
	CommentsCount int       `json:"comments_count"`
	RepostsCount  int       `json:"reposts_count"`
	ViewsCount    int       `json:"views_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Topic         Topic     `db:"-"`
	AuthorName    string    `json:"author_name"`
	AuthorAvatar  string    `json:"author_avatar"`
}
