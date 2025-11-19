package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) CreateAnonymousAppeal(ctx context.Context, appealDTO dto.AppealInputDto) (models.Appeal, error) {
	catID, err := strconv.Atoi(appealDTO.CategoryID)
	if err != nil {
		return models.Appeal{}, errors.New("invalid category_id")
	}

	appeal := models.Appeal{
		AppealID:           uuid.New(),
		CreatorID:          uuid.New(),
		EmailRegistered:    appealDTO.EmailRegistered,
		Status:             models.Status(appealDTO.Status),
		ProblemDescription: appealDTO.ProblemDescription,
		Name:               appealDTO.Name,
		CategoryID:         catID,
		EmailForConnect:    appealDTO.EmailForConnect,
		ScreenshotURL:      appealDTO.ScreenshotURL,
	}

	created, err := u.appealRepo.CreateNewAppeal(ctx, appeal)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, u.handleError(err)
	}

	return created, nil
}
