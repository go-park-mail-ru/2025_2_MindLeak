package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	authmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/mock"
	"github.com/google/uuid"
)

func TestHandler_Logout_Success(t *testing.T) {
	usecase := new(authmock.AuthUsecaseMock)
	sessionID := uuid.New()

	usecase.On("Logout", mock.Anything, sessionID).Return(true, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: sessionID.String(),
	})
	rr := httptest.NewRecorder()

	h.Logout(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "logged out")

	usecase.AssertExpectations(t)
}
