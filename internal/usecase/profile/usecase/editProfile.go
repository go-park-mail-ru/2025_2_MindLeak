package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/security"
	"github.com/google/uuid"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	InvalidPassword = errors.New("invalid password")
	InvalidName     = errors.New("invalid name")
)

func (u *Usecase) EditProfile(ctx context.Context, sessionID uuid.UUID, newProfile models.Profile, newUser models.User) (models.Profile, models.User, error) {
	session, err := u.sessionRepo.GetSessionById(ctx, sessionID)
	if err != nil {
		return models.Profile{}, models.User{}, u.handleError(err)
	}

	userID := session.UserId

	oldProfile, err := u.profileRepo.GetProfile(ctx, userID)
	if err != nil {
		return models.Profile{}, models.User{}, u.handleError(err)
	}

	oldUser, err := u.userRepo.GetUserById(ctx, userID)
	if err != nil {
		return models.Profile{}, models.User{}, u.handleError(err)
	}

	if err := u.ValidateProfileData(newProfile); err != nil {
		return models.Profile{}, models.User{}, u.handleError(err)
	}

	if err := ValidateUserData(newUser); err != nil {
		return models.Profile{}, models.User{}, u.handleError(err)
	}

	profileChanged := u.IsEditedProfileData(newProfile, &oldProfile)
	userChanged, err := u.IsEditedUserData(newUser, &oldUser)
	if err != nil {
		return models.Profile{}, models.User{}, u.handleError(err)
	}

	var updatedProfile models.Profile
	var updatedUser models.User

	if profileChanged {
		updatedProfile, err = u.profileRepo.UpdateProfile(ctx, oldProfile)
		if err != nil {
			return models.Profile{}, models.User{}, u.handleError(err)
		}
	} else {
		updatedProfile = oldProfile
	}

	if userChanged {
		updatedUser, err = u.userRepo.UpdateUser(ctx, oldUser)
		if err != nil {
			return models.Profile{}, models.User{}, u.handleError(err)
		}
	} else {
		updatedUser = oldUser
	}

	return updatedProfile, updatedUser, nil
}

func (u *Usecase) ValidateProfileData(p models.Profile) error {
	now := time.Now()

	if !p.DateOfBirth.IsZero() {
		if p.DateOfBirth.After(now) {
			return fmt.Errorf("date of birth is in the future")
		}
		age := now.Year() - p.DateOfBirth.Year()
		if age > 130 {
			return fmt.Errorf("age cannot exceed 130 years")
		}
	}

	phoneRegex := regexp.MustCompile(`^[0-9+\-\(\) ]{7,20}$`)
	if p.Phone != "" && !phoneRegex.MatchString(p.Phone) {
		return fmt.Errorf("invalid phone format")
	}

	if len(p.Country) > 100 {
		return fmt.Errorf("invalid country")
	}

	if p.Age < 0 || p.Age > 120 {
		return fmt.Errorf("invalid age")
	}

	if len(p.Language) > 50 {
		return fmt.Errorf("invalid language")
	}

	return nil
}

func ValidateUserData(u models.User) error {
	if err := validateName(u.Name); err != nil {
		return err
	}

	if u.Password != "" {
		if err := validatePassword(u.Password); err != nil {
			return err
		}
	}

	return nil
}

func (u *Usecase) IsEditedProfileData(newProfile models.Profile, oldProfile *models.Profile) bool {
	profileChanged := false

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
		cleanDescription := u.sanitizer.Sanitize(newProfile.Description)
		oldProfile.Description = cleanDescription
		profileChanged = true
	}
	if newProfile.Language != "" && newProfile.Language != oldProfile.Language {
		oldProfile.Language = newProfile.Language
		profileChanged = true
	}

	return profileChanged
}

func (u *Usecase) IsEditedUserData(newUser models.User, oldUser *models.User) (bool, error) {
	userChanged := false

	if newUser.Name != "" && newUser.Name != oldUser.Name {
		oldUser.Name = newUser.Name
		userChanged = true
	}
	if newUser.Avatar != "" && newUser.Avatar != oldUser.Avatar {
		oldUser.Avatar = newUser.Avatar
		userChanged = true
	}
	if newUser.Password != "" {
		hashed, err := security.HashPassword(newUser.Password)
		if err != nil {
			return false, err
		}
		hashedStr := fmt.Sprintf("%x", hashed)
		oldUser.Password = hashedStr
		userChanged = true
	}

	return userChanged, nil
}

func validatePassword(password string) error {
	if password == "" {
		return InvalidPassword
	}

	if utf8.RuneCountInString(password) < 4 {
		return InvalidPassword
	}

	if strings.Contains(password, " ") {
		return InvalidPassword
	}

	if utf8.RuneCountInString(password) > 64 {
		return InvalidPassword
	}

	return nil
}

func validateName(name string) error {
	if name == "" {
		return InvalidName
	}

	if strings.Contains(name, " ") {
		return InvalidName
	}

	if utf8.RuneCountInString(name) < 4 {
		return InvalidName
	}

	if utf8.RuneCountInString(name) > 32 {
		return InvalidName
	}

	return nil
}
