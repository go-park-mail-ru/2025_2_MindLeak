package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) GetArticlesByAuthorId(ctx context.Context, sessionID uuid.UUID) ([]models.Article, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, "session not found: %v", err)
		return nil, u.handleError(err)
	}

	articles, err := u.articleRepo.GetArticlesByAuthorId(ctx, session.UserId)
	if err != nil {
		logger.Error(ctx, "get own articles failed: %v", err)
		return nil, u.handleError(err)
	}

	return articles, nil
}
