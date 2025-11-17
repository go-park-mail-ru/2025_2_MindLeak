package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

func (u *Usecase) GetCategories(ctx context.Context) ([]models.AppealCategory, error) {
	categories, err := u.appealRepo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
