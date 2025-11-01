package profile

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

type Usecase interface {
	ShowProfile(ctx context.Context, sessionID uuid.UUID) (models.Profile, error)
	EditProfile(ctx context.Context) (models.Profile, error)
}
