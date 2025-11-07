package categories

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

type Usecase interface {
	GetFeedByTopic(ctx context.Context, topic string, offset int) ([]models.Article, error)
}
