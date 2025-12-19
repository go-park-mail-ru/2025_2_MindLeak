package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/http/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) CreateAppeal(ctx context.Context, sessionID uuid.UUID, appealDTO dto.AppealInputDto) (models.Appeal, error) {

	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, u.handleError(err)
	}

	catID, err := strconv.Atoi(appealDTO.CategoryID)
	if err != nil {
		return models.Appeal{}, errors.New("invalid category_id")
	}

	authorID := session.UserId

	appeal := models.Appeal{
		//AppealID:           uuid.New(),
		CreatorID:          authorID,
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
