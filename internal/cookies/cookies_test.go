package cookies

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCookies(t *testing.T) {
	sessionID := uuid.New()

	tests := []struct {
		name     string
		setupReq func() *http.Request
	}{
		{
			name: "SetCookie sets correct values",
			setupReq: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/", nil)
			},
		},
		{
			name: "GetCookie returns existing cookie",
			setupReq: func() *http.Request {
				r := httptest.NewRequest(http.MethodGet, "/", nil)
				r.AddCookie(&http.Cookie{Name: SessionID, Value: sessionID.String()})
				return r
			},
		},
		{
			name: "GetCookie missing returns error",
			setupReq: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/", nil)
			},
		},
		{
			name: "DeleteCookie removes existing cookie",
			setupReq: func() *http.Request {
				r := httptest.NewRequest(http.MethodGet, "/", nil)
				r.AddCookie(&http.Cookie{Name: SessionID, Value: sessionID.String()})
				return r
			},
		},
		{
			name: "DeleteCookie missing returns error",
			setupReq: func() *http.Request {
				return httptest.NewRequest(http.MethodGet, "/", nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setupReq()
			w := httptest.NewRecorder()

			switch tt.name {
			case "SetCookie sets correct values":
				SetCookie(w, sessionID)
				res := w.Result()
				cookies := res.Cookies()
				assert.Len(t, cookies, 1)

				c := cookies[0]
				assert.Equal(t, SessionID, c.Name)
				assert.Equal(t, sessionID.String(), c.Value)
				assert.True(t, c.HttpOnly)
				assert.False(t, c.Secure)
				assert.Equal(t, "/", c.Path)

			case "GetCookie returns existing cookie":
				c, err := GetCookie(req)
				assert.NoError(t, err)
				assert.NotNil(t, c)
				assert.Equal(t, sessionID.String(), c.Value)

			case "GetCookie missing returns error":
				c, err := GetCookie(req)
				assert.Nil(t, c)
				assert.EqualError(t, err, "cookie not found")

			case "DeleteCookie removes existing cookie":
				err := DeleteCookie(w, req)
				assert.NoError(t, err)

				res := w.Result()
				cookies := res.Cookies()
				assert.Len(t, cookies, 1)

				c := cookies[0]
				assert.Equal(t, SessionID, c.Name)
				assert.Equal(t, -1, c.MaxAge)

			case "DeleteCookie missing returns error":
				err := DeleteCookie(w, req)
				assert.EqualError(t, err, "cookie not found")
			}
		})
	}
}
