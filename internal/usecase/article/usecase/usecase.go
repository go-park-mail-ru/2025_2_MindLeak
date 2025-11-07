package usecase

import (
	"errors"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/minio_client"
)

var (
	ErrArticleExists   = errors.New("article is already exists")
	ErrArticleNotFound = errors.New("article not found")
	ErrServerError     = errors.New("internal apiserver error")
)

type Usecase struct {
	articleRepo article.ArticleRepository
	sessionRepo session.SessionRepository
	minioClient *minio_client.Client
}

func NewArticleUsecase(articleRepo article.ArticleRepository, sessionRepo session.SessionRepository, minioClient *minio_client.Client) *Usecase {
	return &Usecase{
		articleRepo: articleRepo,
		sessionRepo: sessionRepo,
		minioClient: minioClient,
	}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, article.ErrArticleExists):
		return ErrArticleExists
	case errors.Is(err, article.ErrArticleNotFound):
		return ErrArticleNotFound
	default:
		return ErrServerError
	}
}
