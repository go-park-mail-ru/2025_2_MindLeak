package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

func (u *Usecase) GetFeedByTopic(ctx context.Context, topic string, offset int) ([]models.Article, error) {

	articles, err := u.articleRepo.GetArticlesByTopic(ctx, topic, offset)
	if err != nil {
		return nil, u.handleError(err)
	}

	return articles, nil
}
