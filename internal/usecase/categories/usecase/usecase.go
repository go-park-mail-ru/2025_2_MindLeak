package usecase

import "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"

type Usecase struct {
	articleRepo article.ArticleRepository
}

func NewCategoriesUsecase(articleRepo article.ArticleRepository) *Usecase {
	return &Usecase{articleRepo: articleRepo}
}

//func (u Usecase) handleError(err error) error {}
