package mock

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
)

type AuthUsecaseMock struct {
	mock.Mock
}

func (m *AuthUsecaseMock) Registration(ctx context.Context, user models.User) (dto.RegisteredUserDto, uuid.UUID, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(dto.RegisteredUserDto), args.Get(1).(uuid.UUID), args.Error(2)
}

func (m *AuthUsecaseMock) Login(ctx context.Context, user models.User) (dto.RegisteredUserDto, uuid.UUID, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(dto.RegisteredUserDto), args.Get(1).(uuid.UUID), args.Error(2)
}

func (m *AuthUsecaseMock) Logout(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	args := m.Called(ctx, sessionID)
	return args.Bool(0), args.Error(1)
}

func (m *AuthUsecaseMock) Me(ctx context.Context, sessionID uuid.UUID) (dto.RegisteredUserDto, error) {
	args := m.Called(ctx, sessionID)
	return args.Get(0).(dto.RegisteredUserDto), args.Error(1)
}
