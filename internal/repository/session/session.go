package session

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/pkg/logger"
	"github.com/gomodule/redigo/redis"
	"time"

	"github.com/google/uuid"
)

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrCreatingSession = errors.New("failed to create session")
	ErrSettingSession  = errors.New("setting session failed")
)

type SessionRepository interface {
	CreateSession(ctx context.Context) (models.Session, error)
	GetSessionById(ctx context.Context, sessionId uuid.UUID) (models.Session, error)
	SetSessionUserId(ctx context.Context, sessionId uuid.UUID, userId uuid.UUID) (models.Session, error)
	DeleteSessionById(ctx context.Context, sessionId uuid.UUID) (bool, error)
}

type RedisSessionManager struct {
	redisConn redis.Conn
}

func NewRedisSessionManager(conn redis.Conn) *RedisSessionManager {
	return &RedisSessionManager{redisConn: conn}
}

func (s *RedisSessionManager) CreateSession(ctx context.Context) (models.Session, error) {
	sessionID := uuid.New()

	_, err := s.redisConn.Do("SETEX",
		"session:"+sessionID.String(),
		int(24*time.Hour/time.Second),
		"",
	)
	if err != nil {
		logger.Error(ctx, "Error creating empty session: %v", err)
		return models.Session{}, ErrCreatingSession
	}

	logger.Info(ctx, "🆕 Empty session %s created", sessionID)

	return models.Session{
		SessionId: sessionID,
		UserId:    uuid.Nil,
	}, nil

}

func (s *RedisSessionManager) GetSessionById(ctx context.Context, sessionId uuid.UUID) (models.Session, error) {
	value, err := redis.String(s.redisConn.Do("GET", "session:"+sessionId.String()))
	if err == redis.ErrNil {
		logger.Error(ctx, "Session %s not found", sessionId.String())
		return models.Session{}, ErrSessionNotFound
	}
	if err != nil {
		logger.Error(ctx, "Error getting session %s: %v", sessionId.String(), err)
		return models.Session{}, ErrSessionNotFound
	}

	var userID uuid.UUID
	if value != "" {
		userID, err = uuid.Parse(value)
		if err != nil {
			logger.Error(ctx, "Invalid UUID stored in session %s: %v", sessionId.String(), err)
			return models.Session{}, ErrSessionNotFound
		}
	}

	logger.Info(ctx, "Session %s found for user %s", sessionId.String(), userID.String())

	return models.Session{
		SessionId: sessionId,
		UserId:    userID,
	}, nil
}

func (s *RedisSessionManager) SetSessionUserId(ctx context.Context, sessionId uuid.UUID, userId uuid.UUID) (models.Session, error) {
	exists, err := redis.Int(s.redisConn.Do("EXISTS", "session:"+sessionId.String()))
	if err != nil {
		logger.Error(ctx, "Error checking session existence: %v", err)
		return models.Session{}, err
	}
	if exists == 0 {
		logger.Warn(ctx, "Session %s not found when setting user", sessionId.String())
		return models.Session{}, ErrSessionNotFound
	}

	_, err = s.redisConn.Do("SETEX",
		"session:"+sessionId.String(),
		int(24*time.Hour/time.Second),
		userId.String(),
	)
	if err != nil {
		logger.Error(ctx, "Error setting user ID in session %s: %v", sessionId.String(), err)
		return models.Session{}, ErrSettingSession
	}

	logger.Info(ctx, "User %s linked to session %s", userId.String(), sessionId.String())

	return models.Session{
		SessionId: sessionId,
		UserId:    userId,
	}, nil
}

func (s *RedisSessionManager) DeleteSessionById(ctx context.Context, sessionId uuid.UUID) (bool, error) {
	deleted, err := redis.Int(s.redisConn.Do("DEL", "session:"+sessionId.String()))
	if err != nil {
		logger.Error(ctx, "Error deleting session %s: %v", sessionId.String(), err)
		return false, err
	}

	if deleted == 0 {
		logger.Warn(ctx, "Session %s not found for deletion", sessionId.String())
		return false, ErrSessionNotFound
	}

	logger.Info(ctx, "Session %s deleted", sessionId.String())
	return true, nil
}
