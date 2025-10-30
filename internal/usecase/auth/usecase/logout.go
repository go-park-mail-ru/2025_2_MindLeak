package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (u *Usecase) Logout(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	flag, err := u.sessionRepo.DeleteSessionById(ctx, sessionID)
	if err != nil {
		return false, fmt.Errorf("get auth usecase: %w", err)
	}
	return flag, err
}
