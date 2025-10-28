package article

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/article/dto"
)

type Usecase interface {
	Feed(ctx context.Context, feed models.Feed) (dto.ReceivedFeedDTO, error)
}
