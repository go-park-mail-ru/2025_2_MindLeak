package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) GetAppeals(ctx context.Context, sessionID uuid.UUID) ([]models.Appeal, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	authorID := session.UserId

	apeals, err := u.appealRepo.GetMyAppeals(ctx, authorID)
	if err != nil {
		return nil, err
	}

	return apeals, nil
}
