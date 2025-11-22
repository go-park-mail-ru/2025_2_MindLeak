package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) GetSession(ctx context.Context, sessionID uuid.UUID) (models.Session, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return models.Session{}, u.handleError(err)
	}

	return session, nil

}
