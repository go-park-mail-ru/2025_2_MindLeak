package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

func (u *Usecase) SearchArticles(ctx context.Context, queryText string) ([]models.Article, error) {
	articles, err := u.articleRepo.SearchArticles(ctx, queryText)
	if err != nil {
		return nil, u.handleError(err)
	}

	return articles, nil
}
