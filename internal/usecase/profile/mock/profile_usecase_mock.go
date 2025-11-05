package mock

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/dto"
)

type ProfileUsecaseMock struct {
	mock.Mock
}

func (m *ProfileUsecaseMock) ShowProfile(ctx context.Context, targetUserId uuid.UUID) (dto.ProfileDto, error) {
	args := m.Called(ctx, targetUserId)
	return args.Get(0).(dto.ProfileDto), args.Error(1)
}

func (m *ProfileUsecaseMock) EditProfile(ctx context.Context, sessionID uuid.UUID, newProfile models.Profile, newUser models.User) (models.Profile, models.User, error) {
	args := m.Called(ctx, sessionID, newProfile, newUser)
	return args.Get(0).(models.Profile), args.Get(1).(models.User), args.Error(2)
}

func (m *ProfileUsecaseMock) GetSession(ctx context.Context, sessionID uuid.UUID) (models.Session, error) {
	args := m.Called(ctx, sessionID)
	return args.Get(0).(models.Session), args.Error(1)
}

func (m *ProfileUsecaseMock) UploadAvatar(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.User, error) {
	args := m.Called(ctx, userID, file, header)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *ProfileUsecaseMock) DeleteProfile(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	args := m.Called(ctx, sessionID)
	return args.Bool(0), args.Error(1)
}

func (m *ProfileUsecaseMock) DeleteAvatar(ctx context.Context, sessionID uuid.UUID) (models.User, error) {
	args := m.Called(ctx, sessionID)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *ProfileUsecaseMock) UploadCover(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Profile, error) {
	args := m.Called(ctx, userID, file, header)
	return args.Get(0).(models.Profile), args.Error(1)
}

func (m *ProfileUsecaseMock) DeleteCover(ctx context.Context, sessionID uuid.UUID) (models.Profile, error) {
	args := m.Called(ctx, sessionID)
	return args.Get(0).(models.Profile), args.Error(1)
}
