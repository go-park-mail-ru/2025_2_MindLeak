package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
)

func (u *Usecase) ShowTopBlogs(ctx context.Context) ([]models.User, error) {
	users, err := u.subscriptionsRepo.GetTopBlogs(ctx)
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, u.handleError(err)
	}

	return users, nil
}
