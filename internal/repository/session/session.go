package session

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var ErrSessionNotFound = errors.New("session not found")

type SessionRepository interface {
	CreateSession(ctx context.Context) (*Session, error)
	GetSessionById(ctx context.Context, sessionId uuid.UUID) (*Session, error)
	SetSessionUserId(ctx context.Context, sessionId uuid.UUID, userId uuid.UUID) (*Session, error)
	DeleteSessionById(ctx context.Context, sessionId uuid.UUID) (bool, error)
}

type Session struct {
	SessionId uuid.UUID
	UserId    uuid.UUID
}

type InMemorySession struct {
	Sessions map[uuid.UUID]uuid.UUID
	mu       sync.RWMutex
}

func NewInMemorySession() *InMemorySession {
	return &InMemorySession{
		Sessions: make(map[uuid.UUID]uuid.UUID),
	}
}

func (mem *InMemorySession) CreateSession(ctx context.Context) (*Session, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	SessionId := uuid.New()
	Session := &Session{
		UserId:    uuid.UUID{},
		SessionId: SessionId,
	}
	mem.Sessions[SessionId] = Session.UserId
	return Session, nil
}

func (mem *InMemorySession) GetSessionById(ctx context.Context, sessionId uuid.UUID) (*Session, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	if UserId, exists := mem.Sessions[sessionId]; exists {
		session := &Session{
			SessionId: sessionId,
			UserId:    UserId,
		}
		return session, nil
	} else {
		return nil, fmt.Errorf("%w: %s", ErrSessionNotFound, sessionId)
	}
}

func (mem *InMemorySession) SetSessionUserId(ctx context.Context, sessionId uuid.UUID, userId uuid.UUID) (*Session, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	if _, exists := mem.Sessions[sessionId]; exists {
		mem.Sessions[sessionId] = userId
		session := &Session{
			SessionId: sessionId,
			UserId:    userId,
		}
		return session, nil
	} else {
		return nil, fmt.Errorf("%w: %s", ErrSessionNotFound, sessionId)
	}
}

func (mem *InMemorySession) DeleteSessionById(ctx context.Context, sessionId uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	if _, exists := mem.Sessions[sessionId]; exists {
		delete(mem.Sessions, sessionId)
		return true, nil
	} else {
		return false, fmt.Errorf("%w: %s", ErrSessionNotFound, sessionId)
	}
}
