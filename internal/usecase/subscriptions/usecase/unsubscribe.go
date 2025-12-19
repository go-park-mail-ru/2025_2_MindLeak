package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (u *Usecase) Unsubscribe(ctx context.Context, targetID uuid.UUID, sessionID uuid.UUID) (bool, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return false, ErrUnauthorized
	}
	followerID := session.UserId
	if followerID == targetID {
		return false, ErrSelfSubscribe
	}

	_, err = u.userRepo.GetUserById(ctx, targetID)
	if err != nil {
		return false, ErrUserNotFound
	}

	exists, err := u.subsRepo.Exists(ctx, followerID, targetID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, ErrNotSubscribed
	}

	err = u.subsRepo.Unsubscribe(ctx, followerID, targetID)
	if err != nil {
		return false, err
	}

	return true, nil
}
