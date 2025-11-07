package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) GetArticleById(ctx context.Context, articleID uuid.UUID) (models.Article, error) {
	article, err := u.articleRepo.GetArticleById(ctx, articleID)
	if err != nil {
		logger.Error(ctx, "article not found: %v", err)
		return models.Article{}, u.handleError(err)
	}
	return article, nil
}
