package mock

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
)

type TopBlogsUsecaseMock struct {
	mock.Mock
}

func (m *TopBlogsUsecaseMock) ShowTopBlogs(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.User), args.Error(1)
}
