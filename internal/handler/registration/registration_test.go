package registration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/user"
	"github.com/stretchr/testify/assert"
)

func TestRegistrationStatus(t *testing.T) {
	type test struct {
		name       string
		body       string
		wantStatus int
	}

	tests := []test{
		{
			name:       "missing password",
			body:       `{"email":"user@mail.com","name":"user"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid password (too short)",
			body:       `{"email":"user@mail.com","password":"12","name":"user"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid password (contains space)",
			body:       `{"email":"user@mail.com","password":"12 34","name":"user"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "valid password",
			body:       `{"email":"user@mail.com","password":"1234","name":"user"}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "missing name",
			body:       `{"email":"user@mail.com","password":"1234"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid name (too short)",
			body:       `{"email":"user@mail.com","password":"1234","name":"a"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid name (contains space)",
			body:       `{"email":"user@mail.com","password":"1234","name":"bad name"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "valid name",
			body:       `{"email":"user@mail.com","password":"1234","name":"validUser"}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "valid email",
			body:       `{"email":"valid@mail.com","password":"1234","name":"validUser"}`,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "invalid email (no @)",
			body:       `{"email":"invalidmail.com","password":"1234","name":"user"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "invalid email (no domain)",
			body:       `{"email":"invalid@","password":"1234","name":"user"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "missing email",
			body:       `{"password":"1234","name":"user"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "missing all",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "duplicate email",
			body:       `{"email":"dup@mail.com","password":"1234","name":"user"}`,
			wantStatus: http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/registration", bytes.NewBufferString(test.body))
			w := httptest.NewRecorder()

			sessions := session.NewInMemorySession()
			users := user.NewInMemoryUser()

			if test.name == "duplicate email" {
				_, _ = users.CreateUser("dup@mail.com", "1234", "user")
			}

			RegistrationHandler(w, req, sessions, users)

			resp := w.Result()
			assert.Equal(t, test.wantStatus, resp.StatusCode, "status code mismatch in case %s", test.name)
		})
	}
}

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func TestRegistrationResponse(t *testing.T) {
	type test struct {
		name          string
		body          string
		wantStatus    int
		wantError     bool
		wantErrorText string
		wantEmail     string
	}

	tests := []test{
		{
			name:          "missing password",
			body:          `{"email":"user@mail.com","name":"user"}`,
			wantStatus:    http.StatusBadRequest,
			wantError:     true,
			wantErrorText: "email, password and name are required",
		},
		{
			name:          "invalid email",
			body:          `{"email":"invalid","password":"1234","name":"user"}`,
			wantStatus:    http.StatusBadRequest,
			wantError:     true,
			wantErrorText: "email is invalid",
		},
		{
			name:       "valid user",
			body:       `{"email":"ok@mail.com","password":"1234","name":"user"}`,
			wantStatus: http.StatusCreated,
			wantEmail:  "ok@mail.com",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/registration", bytes.NewBufferString(test.body))
			w := httptest.NewRecorder()

			sessions := session.NewInMemorySession()
			users := user.NewInMemoryUser()

			RegistrationHandler(w, req, sessions, users)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, test.wantStatus, resp.StatusCode, "status code mismatch")

			data, _ := io.ReadAll(resp.Body)

			if test.wantError {
				var errResp ErrorResponse
				assert.NoError(t, json.Unmarshal(data, &errResp))
				assert.Equal(t, test.wantErrorText, errResp.Error, "error message mismatch")
			} else if test.wantEmail != "" {
				var userResp UserResponse
				assert.NoError(t, json.Unmarshal(data, &userResp))
				assert.Equal(t, test.wantEmail, userResp.Email, "email mismatch")
			}
		})
	}
}
