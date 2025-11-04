package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
	"mime/multipart"
)

func (u *Usecase) UploadCover(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.Profile, error) {
	url, err := u.minioClient.UploadCover(ctx, userID, file, header)
	if err != nil {
		return models.Profile{}, u.handleError(err)
	}

	profile, err := u.profileRepo.GetProfile(ctx, userID)
	if err != nil {
		return models.Profile{}, u.handleError(err)
	}

	profile.CoverURL = url

	updated, err := u.profileRepo.UpdateProfile(ctx, profile)
	if err != nil {
		return models.Profile{}, u.handleError(err)
	}

	return updated, nil
}

func (u *Usecase) DeleteCover(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	err := u.minioClient.DeleteCover(ctx, userID)
	if err != nil {
		return models.Profile{}, u.handleError(err)
	}

	profile, err := u.profileRepo.GetProfile(ctx, userID)
	if err != nil {
		return models.Profile{}, u.handleError(err)
	}

	profile.CoverURL = u.minioClient.GetDefaultCover()

	updated, err := u.profileRepo.UpdateProfile(ctx, profile)
	if err != nil {
		return models.Profile{}, u.handleError(err)
	}

	return updated, nil
}

func (u *Usecase) UploadAvatar(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (models.User, error) {
	url, err := u.minioClient.UploadAvatar(ctx, userID, file, header)
	if err != nil {
		return models.User{}, u.handleError(err)
	}

	user, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return models.User{}, u.handleError(err)
	}

	user.Avatar = url

	updated, err := u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return models.User{}, u.handleError(err)
	}

	return updated, nil
}

func (u *Usecase) DeleteAvatar(ctx context.Context, userID uuid.UUID) (models.User, error) {
	err := u.minioClient.DeleteAvatar(ctx, userID)
	if err != nil {
		return models.User{}, u.handleError(err)
	}

	user, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return models.User{}, u.handleError(err)
	}

	user.Avatar = u.minioClient.GetDefaultAvatar()

	updated, err := u.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return models.User{}, u.handleError(err)
	}

	return updated, nil
}
