package usecase

import (
	"errors"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/appeal"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"
)

var (
	ErrAppealNotFound = errors.New("appeal not found")
	ErrServerError    = errors.New("internal apiserver error")
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

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, appeal.ErrAppealNotFound):
		return ErrAppealNotFound
	default:
		return ErrServerError
	}
}
