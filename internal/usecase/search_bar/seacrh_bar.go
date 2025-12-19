package search_bar

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

type Usecase interface {
	SearchUsers(ctx context.Context, queryText string) ([]models.User, error)
	SearchArticles(ctx context.Context, queryText string) ([]models.Article, error)
}
