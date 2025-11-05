package profile

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/dto"
	profilemock "github.com/go-park-mail-ru/2025_2_MindLeak/internal/usecase/profile/mock"
)

func TestHandler_ShowProfileHandler_Success(t *testing.T) {
	usecase := new(profilemock.ProfileUsecaseMock)
	userID := uuid.New()

	usecase.On("ShowProfile", mock.Anything, userID).
		Return(dto.ProfileDto{Name: "John", Country: "RU"}, nil)

	h := &Handler{Usecase: usecase}

	req := httptest.NewRequest(http.MethodGet, "/profile/"+userID.String(), nil)
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	req.AddCookie(&http.Cookie{Name: "session_id", Value: userID.String()})
	rr := httptest.NewRecorder()

	h.ShowProfileHandler(rr, req)

	res := rr.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)
	require.Contains(t, rr.Body.String(), "John")

	usecase.AssertExpectations(t)
}
