package usecase

import (
	"errors"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	repoSession "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
	repoUser "github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

var (
	InvalidEmail      = errors.New("invalid email")
	InvalidPassword   = errors.New("invalid password")
	InvalidName       = errors.New("invalid name")
	UserExists        = errors.New("user is already registered")
	UserNotFound      = errors.New("user not found")
	ServerError       = errors.New("internal apiserver error")
	SessionNotFound   = errors.New("session not found")
	SessionNotCreated = errors.New("session not created")
	SessionNotSet     = errors.New("session not set")
	UserNotCreated    = errors.New("user not created")
	UserNotUpdated    = errors.New("user not updated")
	UserNotDeleted    = errors.New("user not deleted")
	UserNotGet        = errors.New("user not get")
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
	case errors.Is(err, repoSession.ErrCreatingSession):
		return SessionNotCreated
	case errors.Is(err, repoSession.ErrSettingSession):
		return SessionNotSet
	case errors.Is(err, repoUser.ErrCreatingUser):
		return UserNotCreated
	case errors.Is(err, repoUser.ErrUserExists):
		return UserExists
	case errors.Is(err, repoUser.ErrDeletingUser):
		return UserNotDeleted
	case errors.Is(err, repoUser.ErrGettingUser):
		return UserNotGet
	case errors.Is(err, repoUser.ErrUpdatingUser):
		return UserNotUpdated
	default:
		return ServerError
	}
}
