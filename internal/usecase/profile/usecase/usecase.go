package usecase

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/profile"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

type Usecase struct {
	sessionRepo session.SessionRepository
	userRepo    user.UserRepository
	profileRepo profile.ProfileRepository
}

func NewProfileUsecase(sessionRepo session.SessionRepository, userRepo user.UserRepository, profileRepo profile.ProfileRepository) *Usecase {
	return &Usecase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}

func (u *Usecase) handleError(err error) error {
	///
	return nil
}
