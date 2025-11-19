package subscriptions

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

type Usecase interface {
	Subscribe(ctx context.Context, targetID uuid.UUID, sessionID uuid.UUID) (bool, error)
	Unsubscribe(ctx context.Context, targetID uuid.UUID, sessionID uuid.UUID) (bool, error)
	GetSubscribers(ctx context.Context, userID uuid.UUID) ([]models.User, error)
	GetSubscriptions(ctx context.Context, userID uuid.UUID) ([]models.User, error)
}
