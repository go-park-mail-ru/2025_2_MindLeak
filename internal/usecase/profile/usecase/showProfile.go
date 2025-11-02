package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/dto"
	"github.com/google/uuid"
)

func (u *Usecase) ShowProfile(ctx context.Context, targetUserID uuid.UUID) (dto.ProfileDto, error) {
	profile, err := u.profileRepo.GetProfile(ctx, targetUserID)
	if err != nil {
		return dto.ProfileDto{}, err
	}

	user, err := u.userRepo.GetUserById(ctx, targetUserID)
	if err != nil {
		return dto.ProfileDto{}, err
	}

	out := dto.ProfileDto{
		Phone:       profile.Phone,
		Country:     profile.Country,
		Language:    profile.Language,
		Sex:         profile.Sex,
		DateOfBirth: profile.DateOfBirth.Format("2006-01-02"),
		Age:         profile.Age,
		CoverURL:    profile.CoverURL,
		CreatedAt:   profile.CreatedAt,
		Name:        user.Name,
		Avatar:      user.Avatar,
	}

	return out, nil
}
