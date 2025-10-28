package auth

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	Registration(ctx context.Context, User models.User) (dto.RegisteredUserDto, uuid.UUID, error)
	Login(ctx context.Context, User models.User) (dto.RegisteredUserDto, uuid.UUID, error)
	Logout(ctx context.Context, sessionID uuid.UUID) (bool, error)
	Me(ctx context.Context, sessionID uuid.UUID) (dto.RegisteredUserDto, error)
}
