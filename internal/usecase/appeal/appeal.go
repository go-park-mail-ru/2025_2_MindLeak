package appeal

import (
	"context"
	"mime/multipart"

	dtoHand "github.com/go-park-mail-ru/2025_2_MindLeak/internal/handler/appeal/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/appeal/dto"
	"github.com/google/uuid"
)

type Usecase interface {
	CreateAppeal(ctx context.Context, sessionID uuid.UUID, appealDTO dtoHand.AppealInputDto) (models.Appeal, error)
	CreateAnonymousAppeal(ctx context.Context, appealDTO dtoHand.AppealInputDto) (models.Appeal, error)
	DeleteAppeal(ctx context.Context, sessionID uuid.UUID, appealID uuid.UUID) (bool, error)
	GetAppealByID(ctx context.Context, sessionID uuid.UUID, appealID uuid.UUID) (models.Appeal, error)
	GetAppeals(ctx context.Context, sessionID uuid.UUID, authorID uuid.UUID) ([]models.Appeal, error)
	GetAppealsStatistics(ctx context.Context, sessionID uuid.UUID) (dto.AppealsStatisticsDTO, error)
	UploadScreenshot(ctx context.Context, sessionID uuid.UUID, appealID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Appeal, error)
	DeleteMedia(ctx context.Context, appealID uuid.UUID) (models.Appeal, error)
}
