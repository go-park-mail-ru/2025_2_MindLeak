package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) GetAppealByID(ctx context.Context, sessionID uuid.UUID, appealID uuid.UUID) (models.Appeal, error) {
	_, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return models.Appeal{}, ErrAppealNotFound
	}

	appeal, err := u.appealRepo.GetMyAppeal(ctx, appealID)
	if err != nil {
		return models.Appeal{}, ErrAppealNotFound
	}

	return appeal, nil
}
