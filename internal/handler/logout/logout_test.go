package logout

import (
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/repository/session"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/cookies"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogoutHandler(t *testing.T) {
	type test struct {
		name       string
		setup      func(*session.InMemorySession, *http.Request)
		wantStatus int
		wantBody   string
	}

	tests := []test{
		{
			name:       "no cookie",
			setup:      func(_ *session.InMemorySession, _ *http.Request) {},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"error":"cookie not found"}`,
		},
		{
			name: "invalid cookie value (not uuid)",
			setup: func(_ *session.InMemorySession, r *http.Request) {
				r.AddCookie(&http.Cookie{Name: cookies.SessionID, Value: "not-a-uuid"})
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "valid uuid but session not found",
			setup: func(_ *session.InMemorySession, r *http.Request) {
				r.AddCookie(&http.Cookie{Name: cookies.SessionID, Value: uuid.NewString()})
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   `{"error":"session not found"}`,
		},
		{
			name: "valid session logout",
			setup: func(sessions *session.InMemorySession, r *http.Request) {
				session, _ := sessions.CreateSession()
				r.AddCookie(&http.Cookie{Name: cookies.SessionID, Value: session.SessionId.String()})
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"logged out"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/logout", nil)
			w := httptest.NewRecorder()

			sessions := session.NewInMemorySession()
			test.setup(sessions, req)

			LogoutHandler(w, req, sessions)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, test.wantStatus, resp.StatusCode, "unexpected status")

			if test.wantBody != "" {
				data, _ := io.ReadAll(resp.Body)
				assert.JSONEq(t, test.wantBody, string(data))
			}
		})
	}
}
