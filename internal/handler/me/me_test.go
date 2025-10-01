package me

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type UserResponse struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func TestMeHandlerStatus(t *testing.T) {
	type test struct {
		name         string
		method       string
		setupSession func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID)
		setupUsers   func() (*user.InMemoryUser, uuid.UUID)
		setCookie    bool
		cookieValue  string
		wantStatus   int
	}

	tests := []test{
		{
			name:   "invalid method",
			method: http.MethodPost,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:  false,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			name:   "no cookie",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:  false,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:   "invalid cookie",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:   true,
			cookieValue: "invalid-uuid",
			wantStatus:  http.StatusUnauthorized,
		},
		{
			name:   "invalid session",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				sessionID := uuid.New()
				_, _ = sessions.CreateSession()
				_, _ = sessions.SetSessionUserId(sessionID, uuid.New())
				return sessions, sessionID
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:   true,
			cookieValue: "",
			wantStatus:  http.StatusUnauthorized,
		},
		{
			name:   "user not found",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				sessionID := uuid.New()
				_, _ = sessions.CreateSession()
				_, _ = sessions.SetSessionUserId(sessionID, uuid.New())
				return sessions, sessionID
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:   true,
			cookieValue: "",
			wantStatus:  http.StatusUnauthorized,
		},
		{
			name:   "success",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				session, err := sessions.CreateSession()
				if err != nil {
					panic(err)
				}
				_, err = sessions.SetSessionUserId(session.SessionId, userID)
				if err != nil {
					panic(err)
				}
				return sessions, session.SessionId
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				users := user.NewInMemoryUser()
				user, err := users.CreateUser("user@mail.com", "123", "TestUser")
				if err != nil {
					panic(err)
				}
				return users, user.Id
			},
			setCookie:   true,
			cookieValue: "",
			wantStatus:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/me", nil)
			w := httptest.NewRecorder()

			users, userID := tt.setupUsers()
			sessions, sessionID := tt.setupSession(userID)

			if tt.setCookie {
				cookieValue := tt.cookieValue
				if tt.cookieValue == "" {
					cookieValue = sessionID.String()
				}
				req.AddCookie(&http.Cookie{
					Name:  cookies.SessionID,
					Value: cookieValue,
				})
			}

			MeHandler(w, req, sessions, users)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode, "status code mismatch in case %s", tt.name)
		})
	}
}

func TestMeHandlerResponse(t *testing.T) {
	type test struct {
		name          string
		method        string
		setupSession  func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID)
		setupUsers    func() (*user.InMemoryUser, uuid.UUID)
		setCookie     bool
		cookieValue   string
		wantStatus    int
		wantError     bool
		wantErrorText string
		wantEmail     string
	}

	tests := []test{
		{
			name:   "invalid method",
			method: http.MethodPost,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:     false,
			wantStatus:    http.StatusMethodNotAllowed,
			wantError:     true,
			wantErrorText: "method not allowed",
		},
		{
			name:   "no cookie",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:     false,
			wantStatus:    http.StatusUnauthorized,
			wantError:     true,
			wantErrorText: "cookie not found",
		},
		{
			name:   "invalid cookie",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				return session.NewInMemorySession(), uuid.New()
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:   true,
			cookieValue: "invalid-uuid",
			wantStatus:  http.StatusUnauthorized,
			wantError:   false,
		},
		{
			name:   "invalid session",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				sessionID := uuid.New()
				_, _ = sessions.CreateSession()
				_, _ = sessions.SetSessionUserId(sessionID, uuid.New())
				return sessions, sessionID
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:   true,
			cookieValue: "",
			wantStatus:  http.StatusUnauthorized,
			wantError:   false,
		},
		{
			name:   "user not found",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				sessionID := uuid.New()
				_, _ = sessions.CreateSession()
				_, _ = sessions.SetSessionUserId(sessionID, uuid.New())
				return sessions, sessionID
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				return user.NewInMemoryUser(), uuid.New()
			},
			setCookie:   true,
			cookieValue: "",
			wantStatus:  http.StatusUnauthorized,
			wantError:   false,
		},
		{
			name:   "success",
			method: http.MethodGet,
			setupSession: func(userID uuid.UUID) (*session.InMemorySession, uuid.UUID) {
				sessions := session.NewInMemorySession()
				session, err := sessions.CreateSession()
				if err != nil {
					panic(err)
				}
				_, err = sessions.SetSessionUserId(session.SessionId, userID)
				if err != nil {
					panic(err)
				}
				return sessions, session.SessionId
			},
			setupUsers: func() (*user.InMemoryUser, uuid.UUID) {
				users := user.NewInMemoryUser()
				user, err := users.CreateUser("user@mail.com", "123", "TestUser")
				if err != nil {
					panic(err)
				}
				return users, user.Id
			},
			setCookie:   true,
			cookieValue: "",
			wantStatus:  http.StatusOK,
			wantEmail:   "user@mail.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/me", nil)
			w := httptest.NewRecorder()

			users, userID := tt.setupUsers()
			sessions, sessionID := tt.setupSession(userID)

			if tt.setCookie {
				cookieValue := tt.cookieValue
				if tt.cookieValue == "" {
					cookieValue = sessionID.String()
				}
				req.AddCookie(&http.Cookie{
					Name:  cookies.SessionID,
					Value: cookieValue,
				})
			}

			// Additional check for success case to ensure user and session are valid
			if tt.name == "success" {
				user, err := users.GetUserById(userID)
				assert.NoError(t, err, "user should exist")
				assert.Equal(t, "user@mail.com", user.Email, "user email mismatch")
				session, err := sessions.GetSessionById(sessionID)
				assert.NoError(t, err, "session should exist")
				assert.Equal(t, userID, session.UserId, "session userID mismatch")
			}

			MeHandler(w, req, sessions, users)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode, "status code mismatch")

			data, _ := io.ReadAll(resp.Body)

			if tt.wantError {
				var errResp ErrorResponse
				assert.NoError(t, json.Unmarshal(data, &errResp), "failed to unmarshal error response")
				assert.Equal(t, tt.wantErrorText, errResp.Error, "error message mismatch")
			} else if tt.wantEmail != "" {
				var userResp UserResponse
				assert.NoError(t, json.Unmarshal(data, &userResp), "failed to unmarshal user response")
				assert.Equal(t, tt.wantEmail, userResp.Email, "email mismatch")
				assert.Equal(t, "TestUser", userResp.Name, "name mismatch")
				assert.Equal(t, "https://sun9-88.userapi.com/s/v1/ig2/P_e5HW2lWX3ZxayBg73NnzbHzyhxFCXtBseRjSrN_NbemNC78OpkeYfJeXcTOXqyR8NhSwizZKqJEq_R8PhQo607.jpg?quality=95&as=32x40,48x60,72x90,108x135,160x200,240x300,360x450,480x600,540x675,640x800,720x900,1080x1350,1280x1600,1440x1800,1620x2025&from=bu&cs=1620x0", userResp.Avatar, "avatar mismatch")
			}
		})
	}
}
