package usecase

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
)

type Processor struct {
	userRepo    user.UserRepository
	sessionRepo session.SessionRepository
}

func NewProcessor(userRepo user.UserRepository, sessionRepo session.SessionRepository) *Processor {
	return &Processor{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}
