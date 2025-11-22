package profile

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/dto"
	"github.com/google/uuid"
	"mime/multipart"
)

type Usecase interface {
	ShowProfile(ctx context.Context, targetUserId uuid.UUID) (dto.ProfileDto, error)
	EditProfile(ctx context.Context, sessionID uuid.UUID, newProfile models.Profile, newUser models.User) (models.Profile, models.User, error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (models.Session, error)
	UploadAvatar(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.User, error)
	DeleteProfile(ctx context.Context, sessionID uuid.UUID) (bool, error)
	DeleteAvatar(ctx context.Context, sessionID uuid.UUID) (models.User, error)
	UploadCover(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Profile, error)
	DeleteCover(ctx context.Context, sessionID uuid.UUID) (models.Profile, error)
}
