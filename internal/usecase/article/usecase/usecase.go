package usecase

import (
	"errors"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
)

var (
	ErrArticleExists   = errors.New("article is already exists")
	ErrArticleNotFound = errors.New("article not found")
	ErrServerError     = errors.New("internal server error")
)

type Usecase struct {
	articleRepo article.ArticleRepository
	sessionRepo session.SessionRepository
}

func NewArticleUsecase(articleRepo article.ArticleRepository, sessionRepo session.SessionRepository) *Usecase {
	return &Usecase{
		articleRepo: articleRepo,
		sessionRepo: sessionRepo,
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
