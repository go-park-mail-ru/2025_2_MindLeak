package appeal

import (
	"context"
	"mime/multipart"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/appeal/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	CreateAppeal(ctx context.Context, sessionID uuid.UUID, appealDTO dto.AppealDTO) (models.Appeal, error)
	DeleteAppeal(ctx context.Context, sessionID uuid.UUID, appealID uuid.UUID) (bool, error)
	GetAppealByID(ctx context.Context, appealID uuid.UUID) (models.Appeal, error)
	GetAppeals(ctx context.Context, sessionID uuid.UUID, authorID uuid.UUID) ([]models.Appeal, error)
	GetAppealsStatistics(ctx context.Context, sessionID uuid.UUID) (dto.AppealsStatisticsDTO, error)
	UploadMedia(ctx context.Context, appealID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Appeal, error)
	DeleteMedia(ctx context.Context, appealID uuid.UUID) (models.Appeal, error)
}
