package profile

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	profilemock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/mock"
)

func TestHandler_DeleteProfileHandler_Success(t *testing.T) {
	usecase := new(profilemock.ProfileUsecaseMock)
	sessionID := uuid.New()

	usecase.On("GetSession", mock.Anything, sessionID).
		Return(models.Session{UserId: sessionID}, nil)
	usecase.On("DeleteProfile", mock.Anything, sessionID).
		Return(true, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodDelete, "/profile/delete", nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: sessionID.String()})
	rr := httptest.NewRecorder()

	h.DeleteProfileHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "logged out")

	usecase.AssertExpectations(t)
}
