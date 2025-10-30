package usecase

import (
	"errors"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	repoSession "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
	repoUser "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

var (
	InvalidEmail    = errors.New("invalid email")
	InvalidPassword = errors.New("invalid password")
	InvalidName     = errors.New("invalid name")
	UserExists      = errors.New("user is already registered")
	UserNotFound    = errors.New("user not found")
	ServerError     = errors.New("internal server error")
	SessionNotFound = errors.New("session not found")
)

type Usecase struct {
	userRepo    user.UserRepository
	sessionRepo session.SessionRepository
}

func NewAuthUsecase(userRepo user.UserRepository, sessionRepo session.SessionRepository) *Usecase {
	return &Usecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, repoUser.ErrUserExists):
		return UserExists
	case errors.Is(err, repoUser.ErrUserNotFound):
		return UserNotFound
	case errors.Is(err, repoSession.ErrSessionNotFound):
		return SessionNotFound
	default:
		return ServerError
	}
}
