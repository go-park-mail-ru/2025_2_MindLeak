package dto

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

type CreateArticleInput struct {
	Title   string    `json:"title" validate:"required,min=1,max=200"`
	Content string    `json:"content" validate:"required,min=1"`
	TopicID uuid.UUID `json:"topic_id" validate:"required"`
}

type UpdateArticleInput struct {
	Title   *string    `json:"title,omitempty"`
	Content *string    `json:"content,omitempty"`
	Status  *string    `json:"status,omitempty"`
	TopicID *uuid.UUID `json:"topic_id,omitempty"`
}

type ArticleOutput struct {
	ID            uuid.UUID    `json:"id"`
	AuthorID      uuid.UUID    `json:"author_id"`
	Title         string       `json:"title"`
	Content       string       `json:"content"`
	MediaURL      string       `json:"media_url,omitempty"`
	TopicID       uuid.UUID    `json:"topic_id"`
	Status        string       `json:"status"`
	CommentsCount int          `json:"comments_count"`
	RepostsCount  int          `json:"reposts_count"`
	ViewsCount    int          `json:"views_count"`
	Topic         models.Topic `db:"-"`
	AuthorName    string       `json:"author_name"`
	AuthorAvatar  string       `json:"author_avatar"`
}
