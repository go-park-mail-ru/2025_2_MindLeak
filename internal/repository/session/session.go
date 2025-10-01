package session

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type SessionRepository interface {
	CreateSession() (*Session, error)
	GetSessionById(sessionId uuid.UUID) (*Session, error)
	SetSessionUserId(sessionId uuid.UUID, userId uuid.UUID) (*Session, error)
	DeleteSessionById(sessionId uuid.UUID) (bool, error)
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

func (mem *InMemorySession) CreateSession() (*Session, error) {
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

func (mem *InMemorySession) GetSessionById(sessionId uuid.UUID) (*Session, error) {
	mem.mu.RLock()
	defer mem.mu.RUnlock()

	if UserId, exists := mem.Sessions[sessionId]; exists {
		session := &Session{
			SessionId: sessionId,
			UserId:    UserId,
		}
		return session, nil
	} else {
		return nil, errors.New("session not found")
	}
}

func (mem *InMemorySession) SetSessionUserId(sessionId uuid.UUID, userId uuid.UUID) (*Session, error) {
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
		return nil, errors.New("session not found")
	}
}

func (mem *InMemorySession) DeleteSessionById(sessionId uuid.UUID) (bool, error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()

	if _, exists := mem.Sessions[sessionId]; exists {
		delete(mem.Sessions, sessionId)
		return true, nil
	} else {
		return false, errors.New("session not found")
	}
}
