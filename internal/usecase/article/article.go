package article

import (
	"context"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	Feed(ctx context.Context, feed models.Feed) (dto.ReceivedFeedDTO, error)
	CreateArticle(ctx context.Context, sessionID uuid.UUID, title, content string, topicID int, file *multipart.FileHeader) (models.Article, error)
	DeleteArticle(ctx context.Context, sessionID, articleID uuid.UUID) (bool, error)
	UpdateArticle(ctx context.Context, sessionID uuid.UUID, articleID uuid.UUID, title, content, status *string, topicID *int, file *multipart.FileHeader) (models.Article, error)
	GetArticleById(ctx context.Context, articleID uuid.UUID) (models.Article, error)
	GetArticlesByAuthorId(ctx context.Context, authorID uuid.UUID) ([]models.Article, error)
	UploadArticleMedia(ctx context.Context, articleID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Article, error)
	DeleteArticleMedia(ctx context.Context, articleID uuid.UUID) (models.Article, error)
}
