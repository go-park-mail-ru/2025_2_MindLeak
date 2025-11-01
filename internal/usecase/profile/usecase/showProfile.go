package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
)

func (u *Usecase) ShowProfile(ctx context.Context, sessionID uuid.UUID) (dto.ProfileDto, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		logger.Error(ctx, "ShowProfile - GetSessionById: %v", err)
		return dto.ProfileDto{}, err
	}
	userID := session.UserId

	user, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		logger.Error(ctx, "ShowProfile - GetUserById: %v", err)
		return dto.ProfileDto{}, err
	}

	profile, err := u.profileRepo.GetProfile(ctx, userID)
	if err != nil {
		logger.Error(ctx, "ShowProfile - GetProfile: %v", err)
		return dto.ProfileDto{}, err
	}

	outDto := dto.ProfileDto{
		User:        &user,
		Phone:       profile.Phone,
		Country:     profile.Country,
		Language:    profile.Language,
		Sex:         profile.Sex,
		DateOfBirth: profile.DateOfBirth,
		Age:         profile.Age,
		CoverURL:    profile.CoverURL,
		CreatedAt:   profile.CreatedAt,
	}

	return outDto, nil

}
