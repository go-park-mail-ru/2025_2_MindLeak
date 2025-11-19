package comment

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	GetCommentsByAuthor(ctx context.Context, authorId uuid.UUID) ([]dto.CommentDto, error)
	GetCommentsByArticle(ctx context.Context, articleId uuid.UUID) ([]dto.CommentDto, error)
	AddComment(ctx context.Context, comment dto.CommentDto) (dto.CommentDto, error)
	DeleteComment(ctx context.Context, commentId uuid.UUID) (bool, error)
	UpdateComment(ctx context.Context, comment dto.CommentDto) (dto.CommentDto, error)
	GetSession(ctx context.Context, sessionId uuid.UUID) (models.Session, error)
}
