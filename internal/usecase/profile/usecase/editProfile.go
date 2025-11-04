package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/google/uuid"
)

func (u *Usecase) EditProfile(ctx context.Context, sessionID uuid.UUID, newProfile models.Profile, newUser models.User) (models.Profile, models.User, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return models.Profile{}, models.User{}, err
	}

	userID := session.UserId

	oldProfile, err := u.profileRepo.GetProfile(ctx, userID)
	if err != nil {
		return models.Profile{}, models.User{}, err
	}

	oldUser, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return models.Profile{}, models.User{}, err
	}

	profileChanged := false
	userChanged := false

	// === PROFILE ===
	if newProfile.Sex != "" && newProfile.Sex != oldProfile.Sex {
		oldProfile.Sex = newProfile.Sex
		profileChanged = true
	}
	if newProfile.CoverURL != "" && newProfile.CoverURL != oldProfile.CoverURL {
		oldProfile.CoverURL = newProfile.CoverURL
		profileChanged = true
	}
	if !newProfile.DateOfBirth.IsZero() && !newProfile.DateOfBirth.Equal(oldProfile.DateOfBirth) {
		oldProfile.DateOfBirth = newProfile.DateOfBirth
		profileChanged = true
	}
	if newProfile.Phone != "" && newProfile.Phone != oldProfile.Phone {
		oldProfile.Phone = newProfile.Phone
		profileChanged = true
	}
	if newProfile.Country != "" && newProfile.Country != oldProfile.Country {
		oldProfile.Country = newProfile.Country
		profileChanged = true
	}
	if newProfile.Age != 0 && newProfile.Age != oldProfile.Age {
		oldProfile.Age = newProfile.Age
		profileChanged = true
	}
	if newProfile.Description != "" && newProfile.Description != oldProfile.Description {
		oldProfile.Description = newProfile.Description
		profileChanged = true
	}
	if newProfile.Language != "" && newProfile.Language != oldProfile.Language {
		oldProfile.Language = newProfile.Language
		profileChanged = true
	}

	// === USER ===
	if newUser.Name != "" && newUser.Name != oldUser.Name {
		oldUser.Name = newUser.Name
		userChanged = true
	}
	if newUser.Avatar != "" && newUser.Avatar != oldUser.Avatar {
		oldUser.Avatar = newUser.Avatar
		userChanged = true
	}

	var updatedProfile models.Profile
	var updatedUser models.User

	if profileChanged {
		updatedProfile, err = u.profileRepo.UpdateProfile(ctx, oldProfile)
		if err != nil {
			return models.Profile{}, models.User{}, err
		}
	} else {
		updatedProfile = oldProfile
	}

	if userChanged {
		updatedUser, err = u.userRepo.UpdateUser(ctx, oldUser)
		if err != nil {
			return models.Profile{}, models.User{}, err
		}
	} else {
		updatedUser = oldUser
	}

	return updatedProfile, updatedUser, nil
}
