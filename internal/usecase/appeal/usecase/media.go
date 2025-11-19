package usecase

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) UploadScreenshot(ctx context.Context, sessionID uuid.UUID, appealID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Appeal, error) {
	appeal, err := u.appealRepo.GetMyAppeal(ctx, appealID)
	if err != nil {
		return models.Appeal{}, err
	}
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return models.Appeal{}, err
	}

	userID := session.UserId
	if appeal.CreatorID != userID {
		return models.Appeal{}, errors.New("access denied")
	}

	url, err := u.minioClient.UploadAppealScreenshot(ctx, userID, file, header)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, err
	}

	updated, err := u.appealRepo.UpdateAppealScreenshot(ctx, appealID, url)
	if err != nil {
		logger.Error(ctx, err.Error())
		return models.Appeal{}, err
	}

	return updated, nil
}
