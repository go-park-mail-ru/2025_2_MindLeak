package dto

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

type ReceivedFeedDTO struct {
	Articles []models.Article
}

type UpdateArticleInput struct {
	ArticleID uuid.UUID
	Title     *string `json:"title,omitempty"`
	Content   *string `json:"content,omitempty"`
	// Image     *FileUpload `json:"image,omitempty"`
	Status *string `json:"status,omitempty"`
}
