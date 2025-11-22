package session_test

import (
	"testing"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T, mem *session.InMemorySession)
	}{
		{
			name: "CreateSession creates new session",
			run: func(t *testing.T, mem *session.InMemorySession) {
				sess, err := mem.CreateSession()
				assert.NoError(t, err)
				assert.NotNil(t, sess)
				assert.NotEqual(t, uuid.Nil, sess.SessionId)

				_, exists := mem.Sessions[sess.SessionId]
				assert.True(t, exists)
			},
		},
		{
			name: "GetSessionById returns existing session",
			run: func(t *testing.T, mem *session.InMemorySession) {
				sess, _ := mem.CreateSession()
				got, err := mem.GetSessionById(sess.SessionId)

				assert.NoError(t, err)
				assert.Equal(t, sess.SessionId, got.SessionId)
				assert.Equal(t, sess.UserId, got.UserId)
			},
		},
		{
			name: "GetSessionById returns error if not found",
			run: func(t *testing.T, mem *session.InMemorySession) {
				_, err := mem.GetSessionById(uuid.New())
				assert.EqualError(t, err, "session not found")
			},
		},
		{
			name: "SetSessionUserId updates existing session",
			run: func(t *testing.T, mem *session.InMemorySession) {
				sess, _ := mem.CreateSession()
				userID := uuid.New()

				updated, err := mem.SetSessionUserId(sess.SessionId, userID)
				assert.NoError(t, err)
				assert.Equal(t, userID, updated.UserId)

				assert.Equal(t, userID, mem.Sessions[sess.SessionId])
			},
		},
		{
			name: "SetSessionUserId returns error if not found",
			run: func(t *testing.T, mem *session.InMemorySession) {
				_, err := mem.SetSessionUserId(uuid.New(), uuid.New())
				assert.EqualError(t, err, "session not found")
			},
		},
		{
			name: "DeleteSessionById deletes existing session",
			run: func(t *testing.T, mem *session.InMemorySession) {
				sess, _ := mem.CreateSession()
				ok, err := mem.DeleteSessionById(sess.SessionId)
				assert.True(t, ok)
				assert.NoError(t, err)

				_, exists := mem.Sessions[sess.SessionId]
				assert.False(t, exists)
			},
		},
		{
			name: "DeleteSessionById returns error if not found",
			run: func(t *testing.T, mem *session.InMemorySession) {
				ok, err := mem.DeleteSessionById(uuid.New())
				assert.False(t, ok)
				assert.EqualError(t, err, "session not found")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mem := session.NewInMemorySession()
			test.run(t, mem)
		})
	}
}
