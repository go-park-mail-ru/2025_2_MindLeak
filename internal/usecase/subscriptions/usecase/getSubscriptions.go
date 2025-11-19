package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) GetSubscriptions(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	//Также в дальнейшем тут поместить логику приватности и тд

	subscriptions, err := u.subsRepo.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, u.handleError(err)
	}

	return subscriptions, nil
}
