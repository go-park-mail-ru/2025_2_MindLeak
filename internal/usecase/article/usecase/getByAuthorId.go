package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) GetArticlesByAuthorID(ctx context.Context, authorID uuid.UUID) ([]models.Article, error) {
	articles, err := u.articleRepo.GetArticlesByAuthorId(ctx, authorID)
	if err != nil {
		logger.Error(ctx, "get articles by author id failed: %v", err)
		return nil, u.handleError(err)
	}
	return articles, nil
}
