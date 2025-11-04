package usecase

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/google/uuid"
)

func (u *Usecase) Me(ctx context.Context, sessionID uuid.UUID) (dto.RegisteredUserDto, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, fmt.Errorf("session not found: %w", err)
	}

	user, err := u.userRepo.GetUserById(ctx, session.UserId)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, fmt.Errorf("user not found: %w", err)
	}

	outDto := dto.RegisteredUserDto{
		Id:     user.Id,
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar,
	}

	return outDto, err
}
