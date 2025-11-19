package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
)

func (u *Usecase) GetArticlesByTopic(ctx context.Context, topic string, offset int) ([]models.Article, error) {

	articles, err := u.articleRepo.GetArticlesByTopic(ctx, topic, offset)
	if err != nil {
		logger.Error(ctx, "get articles by topic failed: topic=%s, offset=%d, err=%v", topic, offset, err)
		return nil, u.handleError(err)
	}

	return articles, nil
}
