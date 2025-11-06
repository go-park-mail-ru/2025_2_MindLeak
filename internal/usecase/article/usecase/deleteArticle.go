package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) DeleteArticle(ctx context.Context, sessionID, articleID uuid.UUID) (bool, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, "session not found: %v", err)
		return false, u.handleError(err)
	}

	authorID := session.UserId

	article, err := u.articleRepo.GetArticleById(ctx, articleID)
	if err != nil {
		logger.Error(ctx, "article not found: %v", err)
		return false, u.handleError(err)
	}

	if article.AuthorID != authorID {
		return false, u.handleError(fmt.Errorf("not article owner"))
	}

	if article.MediaURL != "" {
		u.minioClient.DeleteMedia(ctx, article.ID)
	}

	deleted, err := u.articleRepo.DeleteArticle(ctx, articleID)
	if err != nil {
		logger.Error(ctx, "delete failed: %v", err)
		return false, u.handleError(err)
	}

	return deleted, nil
}
