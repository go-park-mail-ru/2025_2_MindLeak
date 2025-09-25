package repository

import (
	"errors"
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

type InMemory struct {
	Sessions map[uuid.UUID]uuid.UUID
}

func (mem *InMemory) CreateSession() (*Session, error) {
	SessionId := uuid.New()
	Session := &Session{
		UserId:    uuid.UUID{},
		SessionId: SessionId,
	}
	mem.Sessions[SessionId] = Session.UserId
	return Session, nil
}

func (mem *InMemory) GetSessionById(sessionId uuid.UUID) (*Session, error) {
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

func (mem *InMemory) SetSessionUserId(sessionId uuid.UUID, userId uuid.UUID) (*Session, error) {
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

func (mem *InMemory) DeleteSessionById(sessionId uuid.UUID) (bool, error) {
	if _, exists := mem.Sessions[sessionId]; exists {
		delete(mem.Sessions, sessionId)
		return true, nil
	} else {
		return false, errors.New("session not found")
	}
}
