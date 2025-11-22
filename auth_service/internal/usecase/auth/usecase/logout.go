package usecase

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) Logout(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	flag, err := u.sessionRepo.DeleteSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return false, fmt.Errorf("get auth usecase: %w", err)
	}
	return flag, err
}
