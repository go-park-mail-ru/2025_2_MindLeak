package usecase

import (
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/subscriptions"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrSelfSubscribe     = errors.New("cannot subscriptions to yourself")
	ErrAlreadySubscribed = errors.New("already subscribed")
	ErrInternal          = errors.New("internal server error")
	ErrNotSubscribed     = errors.New("not subscribed")
)

type Usecase struct {
	subsRepo    subscriptions.SubscriptionRepository
	sessionRepo session.SessionRepository
	userRepo    user.UserRepository
}

func NewSubscriptionsUsecase(subsRepo subscriptions.SubscriptionRepository, sessionRepo session.SessionRepository, userRepo user.UserRepository) *Usecase {
	return &Usecase{
		subsRepo:    subsRepo,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (u *Usecase) handleError(err error) error {
	switch {
	case errors.Is(err, subscriptions.ErrQueryFailed):
		return ErrInternal
	case errors.Is(err, subscriptions.ErrScanFailed):
		return ErrInternal
	case errors.Is(err, session.ErrSessionNotFound):
		return ErrUnauthorized
	case errors.Is(err, user.ErrUserNotFound):
		return ErrUserNotFound
	default:
		return ErrInternal
	}
}
