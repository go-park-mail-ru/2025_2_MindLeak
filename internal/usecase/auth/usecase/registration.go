package usecase

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/google/uuid"
	"regexp"
	"strings"
	"unicode/utf8"
)

func (u *Usecase) Registration(ctx context.Context, user models.User) (dto.RegisteredUserDto, uuid.UUID, error) {

	if err := validateEmail(user.Email); err != nil {
		return dto.RegisteredUserDto{}, uuid.UUID{}, fmt.Errorf("validate email: %w", err)
	}

	if err := validatePassword(user.Password); err != nil {
		return dto.RegisteredUserDto{}, uuid.UUID{}, fmt.Errorf("validate password: %w", err)
	}

	if err := validateName(user.Name); err != nil {
		return dto.RegisteredUserDto{}, uuid.UUID{}, fmt.Errorf("validate name: %w", err)
	}

	newUser, err := u.userRepo.CreateUser(ctx, user.Email, user.Password, user.Name)
	if err != nil {
		return dto.RegisteredUserDto{}, uuid.UUID{}, fmt.Errorf("create user: %w", err)
	}

	session, err := u.sessionRepo.CreateSession(ctx)
	if err != nil {
		return dto.RegisteredUserDto{}, uuid.UUID{}, fmt.Errorf("create session: %w", err)
	}

	_, err = u.sessionRepo.SetSessionUserId(ctx, session.SessionId, newUser.Id)
	if err != nil {
		return dto.RegisteredUserDto{}, uuid.UUID{}, fmt.Errorf("set session: %w", err)
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
