package usecase

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

type AuthUsecase struct {
	userRepo    user.UserRepository
	sessionRepo session.SessionRepository
}

func NewAuthUsecase(userRepo user.UserRepository, sessionRepo session.SessionRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}
