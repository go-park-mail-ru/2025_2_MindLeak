package usecase

import (
	"errors"
	"github.com/microcosm-cc/bluemonday"

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
	sanitizer   *bluemonday.Policy
}

func NewArticleUsecase(articleRepo article.ArticleRepository, sessionRepo session.SessionRepository, minioClient *minio_client.Client) *Usecase {
	sanitizer := bluemonday.UGCPolicy()
	return &Usecase{
		articleRepo: articleRepo,
		sessionRepo: sessionRepo,
		minioClient: minioClient,
		sanitizer:   sanitizer,
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
