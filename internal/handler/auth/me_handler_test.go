package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	authmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/mock"
	"github.com/google/uuid"
)

func TestHandler_Me_Success(t *testing.T) {
	usecase := new(authmock.AuthUsecaseMock)
	sessionID := uuid.New()

	expectedUser := dto.RegisteredUserDto{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	usecase.On("Me", mock.Anything, sessionID).
		Return(expectedUser, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: sessionID.String(),
	})
	rr := httptest.NewRecorder()

	h.Me(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var got dto.RegisteredUserDto
	err := json.NewDecoder(rr.Body).Decode(&got)
	require.NoError(t, err)
	require.Equal(t, expectedUser.Name, got.Name)
	require.Equal(t, expectedUser.Email, got.Email)

	usecase.AssertExpectations(t)
}
