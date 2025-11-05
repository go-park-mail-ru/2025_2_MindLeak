package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/models"
	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	authmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/mock"
	"github.com/google/uuid"
)

func TestHandler_Login_Success(t *testing.T) {
	usecase := new(authmock.AuthUsecaseMock)

	expectedDto := dto.RegisteredUserDto{Name: "John"}
	expectedUUID := uuid.New()

	usecase.On("Login", mock.Anything, mock.MatchedBy(func(u models.User) bool {
		return u.Email == "user@example.com" && u.Password == "secret"
	})).Return(expectedDto, expectedUUID, nil)

	h := &Handler{Usecase: usecase}

	body := `{"email":"user@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h.Login(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "John")

	usecase.AssertExpectations(t)
}
