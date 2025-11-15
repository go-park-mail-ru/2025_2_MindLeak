package usecase

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/appeal"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"
)

type Usecase struct {
	appealRepo  appeal.AppealRepository
	sessionRepo session.SessionRepository
	minioClient *minio_client.Client
}

func NewAppealUsecase(appealRepo appeal.AppealRepository, sessionRepo session.SessionRepository, minioClient *minio_client.Client) *Usecase {
	return &Usecase{
		appealRepo:  appealRepo,
		sessionRepo: sessionRepo,
		minioClient: minioClient,
	}
}
