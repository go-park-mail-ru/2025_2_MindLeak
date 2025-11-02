package profile

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	ShowProfile(ctx context.Context, targetUserId uuid.UUID) (dto.ProfileDto, error)
	EditProfile(ctx context.Context, sessionID uuid.UUID, newProfile models.Profile, newUser models.User) (models.Profile, models.User, error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (models.Session, error)
}
