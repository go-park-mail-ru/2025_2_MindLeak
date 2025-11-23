package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

func (u *Usecase) SearchUsers(ctx context.Context, queryText string) ([]models.User, error) {
	users, err := u.userRepo.SearchUsers(ctx, queryText)
	if err != nil {
		return nil, u.handleError(err)
	}

	return users, nil
}
