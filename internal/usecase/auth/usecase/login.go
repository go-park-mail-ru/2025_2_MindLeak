package usecase

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/security"
	"github.com/google/uuid"
)

var (
	InvalidCredentials = errors.New("invalid credentials")
)

func (u *Usecase) Login(ctx context.Context, userModel models.User) (dto.RegisteredUserDto, uuid.UUID, error) {

	if userModel.Email == "" || userModel.Password == "" {
		return dto.RegisteredUserDto{}, uuid.UUID{}, InvalidCredentials
	}

	user, err := u.userRepo.GetUserByEmail(ctx, userModel.Email)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	storedPassword, _ := hex.DecodeString(user.Password)

	if !security.CheckPassword(storedPassword, userModel.Password) {
		return dto.RegisteredUserDto{}, uuid.UUID{}, InvalidCredentials
	}

	session, err := u.sessionRepo.CreateSession(ctx)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	_, err = u.sessionRepo.SetSessionUserId(ctx, session.SessionId, user.Id)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return dto.RegisteredUserDto{}, uuid.UUID{}, u.handleError(err)
	}

	outDto := dto.RegisteredUserDto{
		Id:     user.Id,
		Email:  user.Email,
		Name:   user.Name,
		Avatar: user.Avatar,
	}

	return outDto, session.SessionId, nil

}
