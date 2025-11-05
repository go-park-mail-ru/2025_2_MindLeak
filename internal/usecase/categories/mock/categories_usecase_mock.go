package mock

import (
	"context"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/stretchr/testify/mock"
)

type CategoriesUsecaseMock struct {
	mock.Mock
}

func (m *CategoriesUsecaseMock) GetFeedByTopic(ctx context.Context, topic string, offset int) ([]*models.Article, error) {
	args := m.Called(ctx, topic, offset)
	return args.Get(0).([]*models.Article), args.Error(1)
}
