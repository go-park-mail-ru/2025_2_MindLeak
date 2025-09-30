package login

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func TestLoginStatus(t *testing.T) {
	type test struct {
		name       string
		body       string
		setupUsers func() *repository.InMemoryUser
		wantStatus int
	}

	tests := []test{
		{
			name:       "invalid json",
			body:       "{bad json}",
			setupUsers: repository.NewInMemoryUser,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "missing email",
			body:       `{"password":"123"}`,
			setupUsers: repository.NewInMemoryUser,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "missing password",
			body:       `{"email":"user@mail.com"}`,
			setupUsers: repository.NewInMemoryUser,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "user not found",
			body:       `{"email":"ghost@mail.com","password":"123"}`,
			setupUsers: repository.NewInMemoryUser,
			wantStatus: http.StatusNotFound,
		},
		{
			name: "invalid password",
			body: `{"email":"user@mail.com","password":"wrong"}`,
			setupUsers: func() *repository.InMemoryUser {
				users := repository.NewInMemoryUser()
				_, _ = users.CreateUser("user@mail.com", "123", "Test User")
				return users
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "success login",
			body: `{"email":"user@mail.com","password":"123"}`,
			setupUsers: func() *repository.InMemoryUser {
				users := repository.NewInMemoryUser()
				_, _ = users.CreateUser("user@mail.com", "123", "Test User")
				return users
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()

			sessions := session.NewInMemorySession()
			users := tt.setupUsers()

			LoginHandler(w, req, sessions, users)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode, "status code mismatch in case %s", tt.name)
		})
	}
}

func TestLoginResponse(t *testing.T) {
	type test struct {
		name          string
		body          string
		setupUsers    func() *repository.InMemoryUser
		wantStatus    int
		wantError     bool
		wantErrorText string
		wantEmail     string
		checkCookie   bool
	}

	tests := []test{
		{
			name:          "missing email",
			body:          `{"password":"123"}`,
			setupUsers:    repository.NewInMemoryUser,
			wantStatus:    http.StatusBadRequest,
			wantError:     true,
			wantErrorText: "Email or Password is required",
		},
		{
			name:          "user not found",
			body:          `{"email":"ghost@mail.com","password":"123"}`,
			setupUsers:    repository.NewInMemoryUser,
			wantStatus:    http.StatusNotFound,
			wantError:     true,
			wantErrorText: "user not found",
		},
		{
			name: "invalid password",
			body: `{"email":"user@mail.com","password":"wrong"}`,
			setupUsers: func() *repository.InMemoryUser {
				users := repository.NewInMemoryUser()
				_, _ = users.CreateUser("user@mail.com", "123", "Test User")
				return users
			},
			wantStatus:    http.StatusUnauthorized,
			wantError:     true,
			wantErrorText: "invalid password",
		},
		{
			name: "success login",
			body: `{"email":"user@mail.com","password":"123"}`,
			setupUsers: func() *repository.InMemoryUser {
				users := repository.NewInMemoryUser()
				_, _ = users.CreateUser("user@mail.com", "123", "Test User")
				return users
			},
			wantStatus:  http.StatusOK,
			wantEmail:   "user@mail.com",
			checkCookie: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()

			sessions := session.NewInMemorySession()
			users := tt.setupUsers()

			LoginHandler(w, req, sessions, users)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode, "status code mismatch")

			data, _ := io.ReadAll(resp.Body)

			if tt.wantError {
				var errResp ErrorResponse
				assert.NoError(t, json.Unmarshal(data, &errResp))
				assert.Equal(t, tt.wantErrorText, errResp.Error, "error message mismatch")
			} else if tt.wantEmail != "" {
				var userResp UserResponse
				assert.NoError(t, json.Unmarshal(data, &userResp))
				assert.Equal(t, tt.wantEmail, userResp.Email, "email mismatch")

				if tt.checkCookie {
					cookies := resp.Cookies()
					assert.NotEmpty(t, cookies, "cookie must be set on login")
					_, err := uuid.Parse(cookies[0].Value)
					assert.NoError(t, err, "cookie value must be valid UUID")
				}
			}
		})
	}
}
