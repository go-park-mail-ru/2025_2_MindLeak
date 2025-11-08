package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) DeleteProfile(ctx context.Context, userID uuid.UUID) (bool, error) {
	//flag, err := u.profileRepo.DeleteProfile(ctx, userID)
	flag, err := u.userRepo.DeleteUser(ctx, userID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return false, u.handleError(err)
	}

	return flag, nil
}
