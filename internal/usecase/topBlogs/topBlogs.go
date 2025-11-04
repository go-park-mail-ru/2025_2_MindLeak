package topBlogs

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

type Usecase interface {
	ShowTopBlogs(ctx context.Context) ([]models.User, error)
}
