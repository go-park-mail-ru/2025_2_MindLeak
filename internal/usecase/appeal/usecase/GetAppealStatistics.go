package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/appeal/dto"
	"github.com/google/uuid"
)

func (u *Usecase) GetAppealStatistics(ctx context.Context, sessionID uuid.UUID) (dto.AppealsStatisticsDTO, error) {
	_, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return dto.AppealsStatisticsDTO{}, err
	}

	stats, err := u.appealRepo.GetStats(ctx)
	if err != nil {
		return dto.AppealsStatisticsDTO{}, err
	}

	appeals, err := u.appealRepo.GetAllAppeals(ctx)
	if err != nil {
		return dto.AppealsStatisticsDTO{}, err
	}

	appealDTOs := make([]dto.AppealDTO, 0, len(appeals))
	for _, a := range appeals {
		appealDTOs = append(appealDTOs, dto.AppealDTO{
			AppealID:           a.AppealID.String(),
			Status:             string(a.Status),
			CategoryName:       a.Category.Name,
			ProblemDescription: a.ProblemDescription,
			CreatedAt:          a.CreatedAt.Unix(),
		})
	}

	output := dto.AppealsStatisticsDTO{
		Total:      stats.Total,
		ByCategory: stats.ByCategory,
		ByStatus:   stats.ByStatus,
		Appeals:    appealDTOs,
	}

	return output, nil
}
