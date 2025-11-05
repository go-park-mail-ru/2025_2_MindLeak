package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/comment/dto"
)

type CommentUsecaseMock struct {
	mock.Mock
}

func (m *CommentUsecaseMock) GetCommentsByAuthor(ctx context.Context, authorId uuid.UUID) ([]dto.CommentDto, error) {
	args := m.Called(ctx, authorId)
	return args.Get(0).([]dto.CommentDto), args.Error(1)
}

func (m *CommentUsecaseMock) GetCommentsByArticle(ctx context.Context, articleId uuid.UUID) ([]dto.CommentDto, error) {
	args := m.Called(ctx, articleId)
	return args.Get(0).([]dto.CommentDto), args.Error(1)
}

func (m *CommentUsecaseMock) AddComment(ctx context.Context, comment dto.CommentDto) (dto.CommentDto, error) {
	args := m.Called(ctx, comment)
	return args.Get(0).(dto.CommentDto), args.Error(1)
}

func (m *CommentUsecaseMock) DeleteComment(ctx context.Context, commentId uuid.UUID) (bool, error) {
	args := m.Called(ctx, commentId)
	return args.Bool(0), args.Error(1)
}

func (m *CommentUsecaseMock) UpdateComment(ctx context.Context, comment dto.CommentDto) (dto.CommentDto, error) {
	args := m.Called(ctx, comment)
	return args.Get(0).(dto.CommentDto), args.Error(1)
}

func (m *CommentUsecaseMock) GetSession(ctx context.Context, sessionId uuid.UUID) (models.Session, error) {
	args := m.Called(ctx, sessionId)
	return args.Get(0).(models.Session), args.Error(1)
}
