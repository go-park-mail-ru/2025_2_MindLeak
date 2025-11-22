package usecase

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/security"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/validation"
	"github.com/google/uuid"
)

func (u *Usecase) Registration(ctx context.Context, user models.User) (dto.RegisteredUserDto, uuid.UUID, error) {

	if err := validation.ValidateEmail(user.Email); err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	if err := validation.ValidatePassword(user.Password); err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	if err := validation.ValidateName(user.Name); err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	hashed, err := security.HashPassword(user.Password)
	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	hashedStr := fmt.Sprintf("%x", hashed)

	newUser, err := u.userRepo.CreateUser(ctx, user.Email, hashedStr, user.Name)
	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	_, err = u.profileRepo.CreateProfile(ctx, newUser.Id)
	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	session, err := u.sessionRepo.CreateSession(ctx)
	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	_, err = u.sessionRepo.SetSessionUserId(ctx, session.SessionId, newUser.Id)
	if err != nil {
		logger.Error(ctx, err.Error())
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	outDto := dto.RegisteredUserDto{
		Id:     newUser.Id,
		Email:  newUser.Email,
		Name:   newUser.Name,
		Avatar: newUser.Avatar,
	}

	return outDto, session.SessionId, nil

}
