package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/dto"
	authmock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/auth/mock"
	"github.com/google/uuid"
)

func TestHandler_Registration_Success(t *testing.T) {
	usecase := new(authmock.AuthUsecaseMock)

	expectedDto := dto.RegisteredUserDto{Name: "Alice", Email: "alice@example.com"}
	expectedUUID := uuid.New()
	usecase.On("Registration", mock.Anything, mock.AnythingOfType("models.User")).
		Return(expectedDto, expectedUUID, nil)

	h := &Handler{Usecase: usecase}

	body := `{"name":"Alice","email":"alice@example.com","password":"12345"}`
	req := httptest.NewRequest(http.MethodPost, "/registration", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	h.Registration(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusCreated, res.StatusCode)

	var resp map[string]interface{}
	err := json.NewDecoder(rr.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, "Alice", resp["name"])
	require.Equal(t, "alice@example.com", resp["email"])

	usecase.AssertExpectations(t)
}
