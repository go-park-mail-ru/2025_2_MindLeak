package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
)

var (
	ArticleNotFound   = errors.New("article not found")
	ArticleExists     = errors.New("article exists")
	InternalServerErr = errors.New("internal server error")
)

type Usecase struct {
	articleRepo article.ArticleRepository
}

func NewCategoriesUsecase(articleRepo article.ArticleRepository) *Usecase {
	return &Usecase{articleRepo: articleRepo}
}

func (u Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, article.ErrArticleNotFound):
		return ArticleNotFound
	case errors.Is(err, article.ErrArticleExists):
		return ArticleExists
	default:
		return InternalServerErr
	}
}
