package usecase

import (
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/article"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

type Usecase struct {
	userRepo    user.UserRepository
	articleRepo article.ArticleRepository
}

func NewSearchBarUsecase(userRepo user.UserRepository, articleRepo article.ArticleRepository) *Usecase {
	return &Usecase{
		userRepo:    userRepo,
		articleRepo: articleRepo,
	}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, user.ErrUserNotFound):
		return fmt.Errorf("user not found: %w", err)

	case errors.Is(err, user.ErrGettingUser):
		return fmt.Errorf("failed to get user: %w", err)

	case errors.Is(err, user.ErrCreatingUser):
		return fmt.Errorf("failed to create user: %w", err)

	case errors.Is(err, article.ErrArticleNotFound):
		return fmt.Errorf("article not found: %w", err)

	default:
		return fmt.Errorf("internal error: %w", err)
	}
}
