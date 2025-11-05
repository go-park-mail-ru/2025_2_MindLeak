package profile

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	profilemock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/mock"
)

func TestHandler_EditProfileHandler_Success(t *testing.T) {
	usecase := new(profilemock.ProfileUsecaseMock)
	sessionID := uuid.New()

	profile := models.Profile{Phone: "123", Country: "RU"}
	user := models.User{Name: "Alice"}

	usecase.On("EditProfile", mock.Anything, sessionID, mock.Anything, mock.Anything).
		Return(profile, user, nil)

	h := &Handler{Usecase: usecase}

	body := `{"phone":"123","country":"RU","name":"Alice"}`
	req := httptest.NewRequest(http.MethodPut, "/profile/edit", strings.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: sessionID.String()})
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	h.EditProfileHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "Alice")

	usecase.AssertExpectations(t)
}
