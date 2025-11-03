package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/google/uuid"
	"regexp"
	"strings"
	"unicode/utf8"
)

func (u *Usecase) Registration(ctx context.Context, user models.User) (dto.RegisteredUserDto, uuid.UUID, error) {

	if err := validateEmail(user.Email); err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	if err := validatePassword(user.Password); err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	if err := validateName(user.Name); err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	newUser, err := u.userRepo.CreateUser(ctx, user.Email, user.Password, user.Name)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	session, err := u.sessionRepo.CreateSession(ctx)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	_, err = u.sessionRepo.SetSessionUserId(ctx, session.SessionId, newUser.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	outDto := dto.RegisteredUserDto{
		Email:  newUser.Email,
		Name:   newUser.Name,
		Avatar: newUser.Avatar,
	}

	return outDto, session.SessionId, nil

}

var emailRequired = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

func validateEmail(email string) error {
	if email == "" {
		return InvalidEmail
	}

	if !emailRequired.MatchString(email) {
		return InvalidEmail
	}

	if utf8.RuneCountInString(email) > 320 {
		return InvalidEmail
	}

	return nil
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
